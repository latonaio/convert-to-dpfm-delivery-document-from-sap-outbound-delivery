package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Output_Formatter"
	"convert-to-dpfm-delivery-document-from-sap-outbound-delivery/config"
	"convert-to-dpfm-delivery-document-from-sap-outbound-delivery/convert_complementer"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	complementer *convert_complementer.ConvertComplementer
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	complementer *convert_complementer.ConvertComplementer,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:          context.Background(),
		conf:         conf,
		rmq:          rmq,
		complementer: complementer,
	}
}

func (c *DPFMAPICaller) AsyncDeliveryDocumentCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	convertFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, convertFin, log, &errs, input, output)
		case "Item":
			// TODO: 実装
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}

	// 後処理
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-convertFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	return nil
}

func (c *DPFMAPICaller) headerCreate(
	wg *sync.WaitGroup, mtx *sync.Mutex,
	errFin chan error, log *logger.Logger, errs *[]error,
	sdc *dpfm_api_input_reader.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) {
	var err error = nil
	var header *dpfm_api_output_formatter.Header
	defer wg.Done()
	defer func() {
		errFin <- err
	}()

	err = c.complementer.ComplementHeader(sdc, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
		return
	}

	// data_platform_orders_header_dataの更新
	header, err = dpfm_api_output_formatter.ConvertToHeader(sdc)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
		return
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Header: *header,
	}
	return
}

func (c *DPFMAPICaller) itemCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, input *dpfm_api_input_reader.SDC) {
	return
}
