package requests

type AddressID struct {
	DeliveryDocument       string  `json:"DeliveryDocument"`
	AddressID              *string `json:"AddressID"`
	BusinessPartnerName1   *string `json:"BusinessPartnerName1"`
	Country                *string `json:"Country"`
	Region                 *string `json:"Region"`
	StreetName             *string `json:"StreetName"`
	CityName               *string `json:"CityName"`
	PostalCode             *string `json:"PostalCode"`
	CorrespondenceLanguage *string `json:"CorrespondenceLanguage"`
	FaxNumber              *string `json:"FaxNumber"`
	PhoneNumber            *string `json:"PhoneNumber"`
}
