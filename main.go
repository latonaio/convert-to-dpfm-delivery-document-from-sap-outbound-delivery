package main

import (
	"context"
	dpfm_api_caller "convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Caller"
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Output_Formatter"
	"convert-to-dpfm-delivery-document-from-sap-outbound-delivery/config"
	"convert-to-dpfm-delivery-document-from-sap-outbound-delivery/convert_complementer"
	"encoding/json"
	"fmt"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.SessionControlQueue(), conf.RMQ.QueueToSQL(), 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	complementer := convert_complementer.NewConvertComplementer(ctx, conf, rmq)
	caller := dpfm_api_caller.NewDPFMAPICaller(conf, rmq, complementer)

	for msg := range iter {
		start := time.Now()
		err = callProcess(rmq, caller, conf, msg)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
		l.Info("process time %v\n", time.Since(start).Milliseconds())
	}
}
func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func callProcess(rmq *rabbitmq.RabbitmqClient, caller *dpfm_api_caller.DPFMAPICaller, conf *config.Conf, msg rabbitmq.RabbitmqMessage) (err error) {
	l := logger.NewLogger()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("error occurred: %w", e)
			return
		}
	}()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	var input dpfm_api_input_reader.SDC
	var output dpfm_api_output_formatter.SDC

	err = json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		l.Error(err)
		return
	}
	err = json.Unmarshal(msg.Raw(), &output)
	if err != nil {
		l.Error(err)
		return
	}

	accepter := getAccepter(&input)

	errs := caller.AsyncDeliveryDocumentCreates(accepter, &input, &output, l)
	if len(errs) != 0 {
		for _, err := range errs {
			l.Error(err)
		}
		return xerrors.New("cannot created")
	}
	//rmq.Send(conf.RMQ.QueueToSQL()[0], input)
	rmq.Send(conf.RMQ.QueueToResponse(), output)
	l.JsonParseOut(output)

	return nil
}

func getAccepter(input *dpfm_api_input_reader.SDC) []string {
	accepter := input.Accepter
	if len(input.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}
	return accepter
}
