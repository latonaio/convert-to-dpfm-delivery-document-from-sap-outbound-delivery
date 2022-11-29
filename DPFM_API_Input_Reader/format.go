package dpfm_api_input_reader

import (
	"convert-to-dpfm-delivery-document-from-sap-outbound-delivery/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.DeliveryDocument
	return &requests.Header{
		DeliveryDocument:              data.DeliveryDocument,
		DeliveryDocumentType:          data.DeliveryDocumentType,
		DocumentDate:                  data.DocumentDate,
		ActualGoodsMovementDate:       data.ActualGoodsMovementDate,
		ActualDeliveryRoute:           data.ActualDeliveryRoute,
		Shippinglocationtimezone:      data.Shippinglocationtimezone,
		Receivinglocationtimezone:     data.Receivinglocationtimezone,
		ActualGoodsMovementTime:       data.ActualGoodsMovementTime,
		BillingDocumentDate:           data.BillingDocumentDate,
		CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
		ConfirmationTime:              data.ConfirmationTime,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		CustomerGroup:                 data.CustomerGroup,
		DeliveryBlockReason:           data.DeliveryBlockReason,
		DeliveryDate:                  data.DeliveryDate,
		DeliveryDocumentBySupplier:    data.DeliveryDocumentBySupplier,
		DeliveryIsInPlant:             data.DeliveryIsInPlant,
		DeliveryPriority:              data.DeliveryPriority,
		DeliveryTime:                  data.DeliveryTime,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		GoodsIssueTime:                data.GoodsIssueTime,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		HeaderGrossWeight:             data.HeaderGrossWeight,
		HeaderNetWeight:               data.HeaderNetWeight,
		HeaderVolume:                  data.HeaderVolume,
		HeaderVolumeUnit:              data.HeaderVolumeUnit,
		HeaderWeightUnit:              data.HeaderWeightUnit,
		IncotermsClassification:       data.IncotermsClassification,
		IsExportDelivery:              data.IsExportDelivery,
		LastChangeDate:                data.LastChangeDate,
		LoadingDate:                   data.LoadingDate,
		LoadingPoint:                  data.LoadingPoint,
		LoadingTime:                   data.LoadingTime,
		MeansOfTransport:              data.MeansOfTransport,
		OrderCombinationIsAllowed:     data.OrderCombinationIsAllowed,
		OrderID:                       data.OrderID,
		OverallDelivConfStatus:        data.OverallDelivConfStatus,
		OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
		OverallGoodsMovementStatus:    data.OverallGoodsMovementStatus,
		OverallPackingStatus:          data.OverallPackingStatus,
		OverallPickingConfStatus:      data.OverallPickingConfStatus,
		OverallPickingStatus:          data.OverallPickingStatus,
		PickingDate:                   data.PickingDate,
		PickingTime:                   data.PickingTime,
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		ReceivingPlant:                data.ReceivingPlant,
		ShippingCondition:             data.ShippingCondition,
		ShippingPoint:                 data.ShippingPoint,
		ShippingType:                  data.ShippingType,
		ShipToParty:                   data.ShipToParty,
		SoldToParty:                   data.SoldToParty,
		Supplier:                      data.Supplier,
		TransportationGroup:           data.TransportationGroup,
		TransportationPlanningDate:    data.TransportationPlanningDate,
		TransportationPlanningTime:    data.TransportationPlanningTime,
	}
}

func (sdc *SDC) ConvertToHeaderPartner() *requests.HeaderPartner {
	dataDeliveryDocument := sdc.DeliveryDocument
	data := sdc.DeliveryDocument.HeaderPartner
	return &requests.HeaderPartner{
		DeliveryDocument: dataDeliveryDocument.DeliveryDocument,
		PartnerFunction:  data.PartnerFunction,
		SDDocument:       data.SDDocument,
		SDDocumentItem:   data.SDDocumentItem,
		Customer:         data.Customer,
		Supplier:         data.Supplier,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataDeliveryDocument := sdc.DeliveryDocument
	data := sdc.DeliveryDocument.DeliveryDocumentItem
	return &requests.Item{
		DeliveryDocument:               dataDeliveryDocument.DeliveryDocument,
		DeliveryDocumentItem:           data.DeliveryDocumentItem,
		BaseUnit:                       data.BaseUnit,
		ActualDeliveryQuantity:         data.ActualDeliveryQuantity,
		Batch:                          data.Batch,
		BatchBySupplier:                data.BatchBySupplier,
		CostCenter:                     data.CostCenter,
		CreationDate:                   data.CreationDate,
		CreationTime:                   data.CreationTime,
		DeliveryDocumentItemCategory:   data.DeliveryDocumentItemCategory,
		DeliveryDocumentItemText:       data.DeliveryDocumentItemText,
		DeliveryQuantityUnit:           data.DeliveryQuantityUnit,
		DistributionChannel:            data.DistributionChannel,
		Division:                       data.Division,
		GLAccount:                      data.GLAccount,
		GoodsMovementStatus:            data.GoodsMovementStatus,
		GoodsMovementType:              data.GoodsMovementType,
		InternationalArticleNumber:     data.InternationalArticleNumber,
		InventorySpecialStockType:      data.InventorySpecialStockType,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		ItemBillingBlockReason:         data.ItemBillingBlockReason,
		ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
		ItemGdsMvtIncompletionSts:      data.ItemGdsMvtIncompletionSts,
		ItemGrossWeight:                data.ItemGrossWeight,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		ItemIsBillingRelevant:          data.ItemIsBillingRelevant,
		ItemPackingIncompletionStatus:  data.ItemPackingIncompletionStatus,
		ItemPickingIncompletionStatus:  data.ItemPickingIncompletionStatus,
		ItemVolume:                     data.ItemVolume,
		LastChangeDate:                 data.LastChangeDate,
		LoadingGroup:                   data.LoadingGroup,
		Material:                       data.Material,
		MaterialByCustomer:             data.MaterialByCustomer,
		MaterialFreightGroup:           data.MaterialFreightGroup,
		NumberOfSerialNumbers:          data.NumberOfSerialNumbers,
		OrderID:                        data.OrderID,
		OrderItem:                      data.OrderItem,
		OriginalDeliveryQuantity:       data.OriginalDeliveryQuantity,
		PackingStatus:                  data.PackingStatus,
		PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
		PickingConfirmationStatus:      data.PickingConfirmationStatus,
		PickingStatus:                  data.PickingStatus,
		Plant:                          data.Plant,
		ProductAvailabilityDate:        data.ProductAvailabilityDate,
		ProductAvailabilityTime:        data.ProductAvailabilityTime,
		ProfitCenter:                   data.ProfitCenter,
		StorageLocation:                data.StorageLocation,
		TransportationGroup:            data.TransportationGroup,
	}
}

func (sdc *SDC) ConvertToAddressID() *requests.AddressID {
	dataDeliveryDocument := sdc.DeliveryDocument
	data := sdc.DeliveryDocument.HeaderPartner.AddressID
	return &requests.AddressID{
		DeliveryDocument:       dataDeliveryDocument.DeliveryDocument,
		AddressID:              data.AddressID,
		BusinessPartnerName1:   data.BusinessPartnerName1,
		Country:                data.Country,
		Region:                 data.Region,
		StreetName:             data.StreetName,
		CityName:               data.CityName,
		PostalCode:             data.PostalCode,
		CorrespondenceLanguage: data.CorrespondenceLanguage,
		FaxNumber:              data.FaxNumber,
		PhoneNumber:            data.PhoneNumber,
	}
}

func (sdc *SDC) ConvertToItemDocumentFlow() *requests.ItemDocumentFlow {
	dataDeliveryDocument := sdc.DeliveryDocument
	dataDeliveryDocumentItem := sdc.DeliveryDocument.DeliveryDocumentItem
	data := sdc.DeliveryDocument.DeliveryDocumentItem.ItemDocumentFlow
	return &requests.ItemDocumentFlow{
		DeliveryDocument:               dataDeliveryDocument.DeliveryDocument,
		DeliveryDocumentItem:           dataDeliveryDocumentItem.DeliveryDocumentItem,
		Deliveryversion:                data.Deliveryversion,
		PrecedingDocument:              data.PrecedingDocument,
		PrecedingDocumentItem:          data.PrecedingDocumentItem,
		PrecedingDocumentCategory:      data.PrecedingDocumentCategory,
		Subsequentdocument:             data.Subsequentdocument,
		QuantityInBaseUnit:             data.QuantityInBaseUnit,
		SubsequentDocumentItem:         data.SubsequentDocumentItem,
		SDFulfillmentCalculationRule:   data.SDFulfillmentCalculationRule,
		SubsequentDocumentCategory:     data.SubsequentDocumentCategory,
		TransferOrderInWrhsMgmtIsConfd: data.TransferOrderInWrhsMgmtIsConfd,
	}
}
