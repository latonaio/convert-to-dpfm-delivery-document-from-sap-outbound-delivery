package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey      string           `json:"connection_key"`
	Result             bool             `json:"result"`
	RedisKey           string           `json:"redis_key"`
	Filepath           string           `json:"filepath"`
	APIStatusCode      int              `json:"api_status_code"`
	RuntimeSessionID   string           `json:"runtime_session_id"`
	BusinessPartner    int              `json:"business_partner"`
	ServiceLabel       string           `json:"service_label"`
	DeliveryDocument   DeliveryDocument `json:"DeliveryDocument"`
	APISchema          string           `json:"api_schema"`
	Accepter           []string         `json:"accepter"`
	DeliveryDocumentID string           `json:"delivery_document"`
	Deleted            bool             `json:"deleted"`
}

type DeliveryDocument struct {
	DeliveryDocument              string               `json:"DeliveryDocument"`
	DeliveryDocumentType          *string              `json:"DeliveryDocumentType"`
	DocumentDate                  *string              `json:"DocumentDate"`
	ActualGoodsMovementDate       *string              `json:"ActualGoodsMovementDate"`
	ActualDeliveryRoute           *string              `json:"ActualDeliveryRoute"`
	Shippinglocationtimezone      *string              `json:"Shippinglocationtimezone"`
	Receivinglocationtimezone     *string              `json:"Receivinglocationtimezone"`
	ActualGoodsMovementTime       *string              `json:"ActualGoodsMovementTime"`
	BillingDocumentDate           *string              `json:"BillingDocumentDate"`
	CompleteDeliveryIsDefined     *bool                `json:"CompleteDeliveryIsDefined"`
	ConfirmationTime              *string              `json:"ConfirmationTime"`
	CreationDate                  *string              `json:"CreationDate"`
	CreationTime                  *string              `json:"CreationTime"`
	CustomerGroup                 *string              `json:"CustomerGroup"`
	DeliveryBlockReason           *string              `json:"DeliveryBlockReason"`
	DeliveryDate                  *string              `json:"DeliveryDate"`
	DeliveryDocumentBySupplier    *string              `json:"DeliveryDocumentBySupplier"`
	DeliveryIsInPlant             *bool                `json:"DeliveryIsInPlant"`
	DeliveryPriority              *string              `json:"DeliveryPriority"`
	DeliveryTime                  *string              `json:"DeliveryTime"`
	GoodsIssueOrReceiptSlipNumber *string              `json:"GoodsIssueOrReceiptSlipNumber"`
	GoodsIssueTime                *string              `json:"GoodsIssueTime"`
	HeaderBillingBlockReason      *string              `json:"HeaderBillingBlockReason"`
	HeaderGrossWeight             *string              `json:"HeaderGrossWeight"`
	HeaderNetWeight               *string              `json:"HeaderNetWeight"`
	HeaderVolume                  *string              `json:"HeaderVolume"`
	HeaderVolumeUnit              *string              `json:"HeaderVolumeUnit"`
	HeaderWeightUnit              *string              `json:"HeaderWeightUnit"`
	IncotermsClassification       *string              `json:"IncotermsClassification"`
	IsExportDelivery              *string              `json:"IsExportDelivery"`
	LastChangeDate                *string              `json:"LastChangeDate"`
	LoadingDate                   *string              `json:"LoadingDate"`
	LoadingPoint                  *string              `json:"LoadingPoint"`
	LoadingTime                   *string              `json:"LoadingTime"`
	MeansOfTransport              *string              `json:"MeansOfTransport"`
	OrderCombinationIsAllowed     *bool                `json:"OrderCombinationIsAllowed"`
	OrderID                       *string              `json:"OrderID"`
	OverallDelivConfStatus        *string              `json:"OverallDelivConfStatus"`
	OverallDelivReltdBillgStatus  *string              `json:"OverallDelivReltdBillgStatus"`
	OverallGoodsMovementStatus    *string              `json:"OverallGoodsMovementStatus"`
	OverallPackingStatus          *string              `json:"OverallPackingStatus"`
	OverallPickingConfStatus      *string              `json:"OverallPickingConfStatus"`
	OverallPickingStatus          *string              `json:"OverallPickingStatus"`
	PickingDate                   *string              `json:"PickingDate"`
	PickingTime                   *string              `json:"PickingTime"`
	PlannedGoodsIssueDate         *string              `json:"PlannedGoodsIssueDate"`
	ReceivingPlant                *string              `json:"ReceivingPlant"`
	ShippingCondition             *string              `json:"ShippingCondition"`
	ShippingPoint                 *string              `json:"ShippingPoint"`
	ShippingType                  *string              `json:"ShippingType"`
	ShipToParty                   *string              `json:"ShipToParty"`
	SoldToParty                   *string              `json:"SoldToParty"`
	Supplier                      *string              `json:"Supplier"`
	TransportationGroup           *string              `json:"TransportationGroup"`
	TransportationPlanningDate    *string              `json:"TransportationPlanningDate"`
	TransportationPlanningTime    *string              `json:"TransportationPlanningTime"`
	HeaderPartner                 HeaderPartner        `json:"HeaderPartner"`
	DeliveryDocumentItem          DeliveryDocumentItem `json:"DeliveryDocumentItem"`
}

type HeaderPartner struct {
	PartnerFunction string    `json:"PartnerFunction"`
	SDDocument      string    `json:"SDDocument"`
	SDDocumentItem  *string   `json:"SDDocumentItem"`
	Customer        *string   `json:"Customer"`
	Supplier        *string   `json:"Supplier"`
	AddressID       AddressID `json:"AddressID"`
}

type AddressID struct {
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

type DeliveryDocumentItem struct {
	DeliveryDocumentItem           string           `json:"DeliveryDocumentItem"`
	BaseUnit                       *string          `json:"BaseUnit"`
	ActualDeliveryQuantity         *string          `json:"ActualDeliveryQuantity"`
	Batch                          *string          `json:"Batch"`
	BatchBySupplier                *string          `json:"BatchBySupplier"`
	CostCenter                     *string          `json:"CostCenter"`
	CreationDate                   *string          `json:"CreationDate"`
	CreationTime                   *string          `json:"CreationTime"`
	DeliveryDocumentItemCategory   *string          `json:"DeliveryDocumentItemCategory"`
	DeliveryDocumentItemText       *string          `json:"DeliveryDocumentItemText"`
	DeliveryQuantityUnit           *string          `json:"DeliveryQuantityUnit"`
	DistributionChannel            *string          `json:"DistributionChannel"`
	Division                       *string          `json:"Division"`
	GLAccount                      *string          `json:"GLAccount"`
	GoodsMovementStatus            *string          `json:"GoodsMovementStatus"`
	GoodsMovementType              *string          `json:"GoodsMovementType"`
	InternationalArticleNumber     *string          `json:"InternationalArticleNumber"`
	InventorySpecialStockType      *string          `json:"InventorySpecialStockType"`
	IsCompletelyDelivered          *bool            `json:"IsCompletelyDelivered"`
	ItemBillingBlockReason         *string          `json:"ItemBillingBlockReason"`
	ItemDeliveryIncompletionStatus *string          `json:"ItemDeliveryIncompletionStatus"`
	ItemGdsMvtIncompletionSts      *string          `json:"ItemGdsMvtIncompletionSts"`
	ItemGrossWeight                *string          `json:"ItemGrossWeight"`
	ItemNetWeight                  *string          `json:"ItemNetWeight"`
	ItemWeightUnit                 *string          `json:"ItemWeightUnit"`
	ItemIsBillingRelevant          *string          `json:"ItemIsBillingRelevant"`
	ItemPackingIncompletionStatus  *string          `json:"ItemPackingIncompletionStatus"`
	ItemPickingIncompletionStatus  *string          `json:"ItemPickingIncompletionStatus"`
	ItemVolume                     *string          `json:"ItemVolume"`
	LastChangeDate                 *string          `json:"LastChangeDate"`
	LoadingGroup                   *string          `json:"LoadingGroup"`
	Material                       *string          `json:"Material"`
	MaterialByCustomer             *string          `json:"MaterialByCustomer"`
	MaterialFreightGroup           *string          `json:"MaterialFreightGroup"`
	NumberOfSerialNumbers          *int             `json:"NumberOfSerialNumbers"`
	OrderID                        *string          `json:"OrderID"`
	OrderItem                      *string          `json:"OrderItem"`
	OriginalDeliveryQuantity       *string          `json:"OriginalDeliveryQuantity"`
	PackingStatus                  *string          `json:"PackingStatus"`
	PartialDeliveryIsAllowed       *string          `json:"PartialDeliveryIsAllowed"`
	PickingConfirmationStatus      *string          `json:"PickingConfirmationStatus"`
	PickingStatus                  *string          `json:"PickingStatus"`
	Plant                          *string          `json:"Plant"`
	ProductAvailabilityDate        *string          `json:"ProductAvailabilityDate"`
	ProductAvailabilityTime        *string          `json:"ProductAvailabilityTime"`
	ProfitCenter                   *string          `json:"ProfitCenter"`
	StorageLocation                *string          `json:"StorageLocation"`
	TransportationGroup            *string          `json:"TransportationGroup"`
	ItemDocumentFlow               ItemDocumentFlow `json:"ItemDocumentFlow"`
}

type ItemDocumentFlow struct {
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
