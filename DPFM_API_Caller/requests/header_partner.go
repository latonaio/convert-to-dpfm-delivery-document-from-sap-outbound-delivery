package requests

type HeaderPartner struct {
	DeliveryDocument string  `json:"DeliveryDocument"`
	PartnerFunction  string  `json:"PartnerFunction"`
	SDDocument       string  `json:"SDDocument"`
	SDDocumentItem   *string `json:"SDDocumentItem"`
	Customer         *string `json:"Customer"`
	Supplier         *string `json:"Supplier"`
}
