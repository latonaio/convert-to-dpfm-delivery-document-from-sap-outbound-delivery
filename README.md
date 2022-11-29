# convert-to-dpfm-delivery-document-from-sap-outbound-delivery

convert-to-dpfm-delivery-document-from-sap-outbound-delivery は、周辺業務システム　を データ連携基盤 と統合することを目的に、SAP 出荷データを データ連携基盤 入出荷データに変換するマイクロサービスです。  
https://xxx.xxx.io/api/FUNCTION_DELIVERY_DOCUMENT_SRV/creates/

## 動作環境

convert-to-dpfm-delivery-document-from-sap-outbound-delivery の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
convert-to-dpfm-delivery-document-from-sap-outbound-delivery が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/FUNCTION_DELIVERY_DOCUMENT_SRV/creates/

## 本レポジトリ に 含まれる API名
convert-to-dpfm-delivery-document-from-sap-outbound-delivery には、次の API をコールするためのリソースが含まれています。  

* A_Header（入出荷 - ヘッダデータ）
* A_HeaderPDF（入出荷 - ヘッダPDFデータ）
* A_HeaderPartner（入出荷 - ヘッダ取引先データ）
* A_HeaderPartnerPlant（入出荷 - ヘッダ取引先プラントデータ）
* A_HeaderPartnerContact（入出荷 - ヘッダ取引先コンタクトデータ）
* A_Item（入出荷 - 明細データ）
* A_ItemPartner（入出荷 - 明細取引先データ）
* A_ItemPartnerPlant（入出荷 - 明細取引先プラントデータ）
* A_Address（入出荷 - 住所データ）

## API への 値入力条件 の 初期値
convert-to-dpfm-delivery-document-from-sap-outbound-delivery において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMDeliveryDocumentCreates",
	"accepter": ["Header"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMDeliveryDocumentCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncDeliveryDocumentCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,

	log *logger.Logger,

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
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}

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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は SAP 出荷データが データ連携基盤 入出荷データ に変換された結果の JSON の例です。  
以下の項目のうち、"DeliveryDocument" ～ "OverallDelivReltdBillgStatus" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。  

```
{
  "connection_key": "response",
  "result": true,
  "redis_key": "abcdefg",
  "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
  "api_status_code": 200,
  "runtime_session_id": "",
  "business_partner": 201,
  "service_label": "FUNCTION_DELIVERY_DOCUMENT",
  "message": {
    "Header": {
      "DeliveryDocument": null,
      "Buyer": null,
      "Seller": null,
      "ReferenceDocument": null,
      "ReferenceDocumentItem": null,
      "OrderID": null,
      "OrderItem": null,
      "ContractType": null,
      "OrderVaridityStartDate": null,
      "OrderValidityEndDate": null,
      "InvoiceScheduleStartDate": null,
      "InvoiceScheduleEndDate": null,
      "IssuingLocationTimeZone": null,
      "ReceivingLocationTimeZone": null,
      "DocumentDate": "2022-10-20",
      "PlannedGoodsIssueDate": "2022-12-20",
      "PlannedGoodsIssueTime": null,
      "PlannedGoodsReceiptDate": null,
      "BillingDocumentDate": "2022-12-30",
      "CompleteDeliveryIsDefined": null,
      "OverallDeliveryStatus": null,
      "CreationDate": null,
      "CreationTime": null,
      "HeaderDeliveryBlockStatus": null,
      "HeaderIssuingBlockStatus": null,
      "HeaderReceivingBlockStatus": null,
      "GoodsIssueOrReceiptSlipNumber": "100000",
      "HeaderBillingStatus": null,
      "HeaderBillingConfStatus": null,
      "HeaderBillingBlockReason": null,
      "HeaderGrossWeight": 500,
      "HeaderNetWeight": 500,
      "HeaderWeightUnit": "10",
      "Incoterms": "CIF",
      "BillToCountry": null,
      "BillFromCountry": null,
      "IsExportImportDelivery": null,
      "LastChangeDate": null,
      "IssuingPlantBusinessPartner": null,
      "IssuingPlant": null,
      "ReceivingPlantBusinessPartner": null,
      "ReceivingPlant": "AB01",
      "DeliverToParty": null,
      "DeliverFromParty": null,
      "TransactionCurrency": null,
      "OverallDelivReltdBillgStatus": null
    },
    "HeaderPartner": null,
    "HeaderPartnerPlant": null
  },
  "api_schema": "SAPOutboundDeliveryReads",
  "accepter": [
    "Header"
  ],
  "delivery_document": "80000000",
  "deleted": false
}
```

