package requests

type ItemDocumentFlow struct {
	DeliveryDocument               string  `json:"DeliveryDocument"`
	DeliveryDocumentItem           string  `json:"DeliveryDocumentItem"`
	Deliveryversion                string  `json:"Deliveryversion"`
	PrecedingDocument              string  `json:"PrecedingDocument"`
	PrecedingDocumentItem          string  `json:"PrecedingDocumentItem"`
	PrecedingDocumentCategory      *string `json:"PrecedingDocumentCategory"`
	Subsequentdocument             string  `json:"Subsequentdocument"`
	QuantityInBaseUnit             *string `json:"QuantityInBaseUnit"`
	SubsequentDocumentItem         *string `json:"SubsequentDocumentItem"`
	SDFulfillmentCalculationRule   *string `json:"SDFulfillmentCalculationRule"`
	SubsequentDocumentCategory     *string `json:"SubsequentDocumentCategory"`
	TransferOrderInWrhsMgmtIsConfd *bool   `json:"TransferOrderInWrhsMgmtIsConfd"`
}
