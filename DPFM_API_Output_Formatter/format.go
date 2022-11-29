package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Input_Reader"
	"fmt"
	"strconv"
)

func ConvertToHeader(
	sdc *dpfm_api_input_reader.SDC,
) (*Header, error) {
	data := sdc.DeliveryDocument
	var err error

	var headerGrossWeightloat32 *float32
	if data.HeaderGrossWeight != nil {
		headerGrossWeightloat32, err = parseFolat32Ptr(*data.HeaderGrossWeight)
		if err != nil {
			return nil, err
		}
	} else {
		headerGrossWeightloat32 = nil
	}

	var headerNetWeightoat32 *float32
	if data.HeaderNetWeight != nil {
		headerNetWeightoat32, err = parseFolat32Ptr(*data.HeaderNetWeight)
		if err != nil {
			return nil, err
		}
	} else {
		headerNetWeightoat32 = nil
	}

	header := Header{
		// DeliveryDocument:              data.DeliveryDocument,
		// Buyer:                         data.Buyer,
		// Seller:                        data.Seller,
		// ReferenceDocument:             data.ReferenceDocument,
		// ReferenceDocumentItem:         data.ReferenceDocumentItem,
		// OrderID:                       data.OrderID,
		// OrderItem:                     data.OrderItem,
		// ContractType:                  data.ContractType,
		// OrderValidityStartDate:        data.OrderValidityStartDate,
		// OrderValidityEndDate:          data.OrderValidityEndDate,
		// InvoiceScheduleStartDate:      data.InvoiceScheduleStartDate,
		// InvoiceScheduleEndDate:        data.InvoiceScheduleEndDate,
		// IssuingPlantTimeZone:          data.IssuingPlantTimeZone,
		// ReceivingPlantTimeZone:        data.ReceivingPlantTimeZone,
		DocumentDate:          data.DocumentDate,
		PlannedGoodsIssueDate: data.PlannedGoodsIssueDate,
		// PlannedGoodsIssueTime:         data.PlannedGoodsIssueTime,
		// PlannedGoodsReceiptDate:       data.PlannedGoodsReceiptDate,
		// PlannedGoodsReceiptTime:       data.PlannedGoodsReceiptTime,
		BillingDocumentDate: data.BillingDocumentDate,
		// CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
		// OverallDeliveryStatus:         data.OverallDeliveryStatus,
		// CreationDate:                  data.CreationDate,
		// CreationTime:                  data.CreationTime,
		// IssuingBlockReason:            data.IssuingBlockReason,
		// ReceivingBlockReason:          data.ReceivingBlockReason,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		// HeaderBillingStatus:           data.HeaderBillingStatus,
		// HeaderBillingConfStatus:       data.HeaderBillingConfStatus,
		// HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		HeaderGrossWeight: headerGrossWeightloat32,
		HeaderNetWeight:   headerNetWeightoat32,
		// HeaderVolume:                  data.HeaderVolume,
		// HeaderVolumeUnit:              data.HeaderVolumeUnit,
		HeaderWeightUnit: data.HeaderWeightUnit,
		Incoterms:        data.IncotermsClassification,
		// IsExportImportDelivery:        data.IsExportImportDelivery,
		// LastChangeDate:                data.LastChangeDate,
		// IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
		// IssuingPlant:                  data.IssuingPlant,
		ReceivingPlant: data.ReceivingPlant,
		// ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
		// DeliverToParty:                data.DeliverToParty,
		// DeliverFromParty:              data.DeliverFromParty,
		// TransactionCurrency: data.TransactionCurrency,
		// OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
	}

	return &header, nil
}

func parseFolat32Ptr(s string) (*float32, error) {
	resFloat64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	resFloat32 := getFloat32Ptr(float32(resFloat64))

	return resFloat32, nil
}

func getFloat32Ptr(f float32) *float32 {
	return &f
}
