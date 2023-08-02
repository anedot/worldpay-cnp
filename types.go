package worldpay

import (
	"encoding/xml"
	"io"
	"net/http"
	"sync"
)

type (
	Client struct {
		Client     *http.Client
		Login      string
		Password   string
		ApiBase    string
		MerchantId string
		Log        io.Writer
		mu         sync.Mutex
	}

	LitleOnlineRequest struct {
		XMLName        xml.Name       `xml:"litleOnlineRequest"`
		Version        string         `xml:"version,attr"`
		XmlNamespace   string         `xml:"xmlns,attr"`
		MerchantId     string         `xml:"merchantId,attr"`
		Authentication Authentication `xml:"authentication"`
		Authorization  *Authorization `xml:"authorization"`
		Capture        *Capture       `xml:"capture"`
		Credit         *Credit        `xml:"credit"`
		EcheckCredit   *EcheckCredit  `xml:"echeckCredit"`
		EcheckSale     *EcheckSale    `xml:"echeckSale"`
		EcheckVoid     *EcheckVoid    `xml:"echeckVoid"`
		Sale           *Sale          `xml:"sale"`
		Void           *Void          `xml:"void"`
	}

	LitleOnlineResponse struct {
		XMLName               xml.Name               `xml:"litleOnlineResponse"`
		Version               string                 `xml:"version,attr"`
		XmlNS                 string                 `xml:"xmlns,attr"`
		Response              string                 `xml:"response,attr"`
		Message               string                 `xml:"message,attr"`
		AuthorizationResponse *AuthorizationResponse `xml:"authorizationResponse,omitempty"`
		CaptureResponse       *CaptureResponse       `xml:"captureResponse,omitempty"`
		CreditResponse        *CreditResponse        `xml:"creditResponse,omitempty"`
		EcheckCreditResponse  *EcheckCreditResponse  `xml:"echeckCreditResponse,omitempty"`
		EcheckSaleResponse    *EcheckSaleResponse    `xml:"echeckSalesResponse,omitempty"`
		EcheckVoidResponse    *EcheckVoidResponse    `xml:"echeckVoidResponse,omitempty"`
		SaleResponse          *SaleResponse          `xml:"saleResponse,omitempty"`
		VoidResponse          *VoidResponse          `xml:"voidResponse,omitempty"`
	}

	Authentication struct {
		User     string `xml:"user"`
		Password string `xml:"password"`
	}

	Authorization struct {
		XMLName                  xml.Name                  `xml:"authorization"`
		Id                       string                    `xml:"id,attr"`
		ReportGroup              string                    `xml:"reportGroup,attr"`
		CustomerId               string                    `xml:"customerId,attr"`
		OrderId                  string                    `xml:"orderId"`
		Amount                   int                       `xml:"amount"`
		OrderSource              string                    `xml:"orderSource"`
		BillToAddress            Address                   `xml:"billToAddress"`
		Card                     Card                      `xml:"card"`
		CardholderAuthentication *CardholderAuthentication `xml:"cardholderAuthentication"`
	}

	Capture struct {
		XMLName      xml.Name      `xml:"capture"`
		Id           string        `xml:"id,attr"`
		ReportGroup  string        `xml:"reportGroup,attr"`
		CustomerId   string        `xml:"customerId,attr"`
		Partial      bool          `xml:"partial,attr"`
		LitleTxnId   string        `xml:"litleTxnId"`
		Amount       int           `xml:"amount"`
		EnhancedData *EnhancedData `xml:"enhancedData"`
	}

	Credit struct {
		XMLName     xml.Name `xml:"credit"`
		Id          string   `xml:"id,attr"`
		ReportGroup string   `xml:"reportGroup,attr"`
		CustomerId  string   `xml:"customerId,attr"`
		LitleTxnId  string   `xml:"litleTxnId"`
		Amount      *int     `xml:"amount,omitempty"`
	}

	EcheckSale struct {
		XMLName       xml.Name `xml:"echeckSale"`
		Id            string   `xml:"id,attr"`
		ReportGroup   string   `xml:"reportGroup,attr"`
		CustomerId    string   `xml:"customerId,attr"`
		OrderId       string   `xml:"orderId"`
		Verify        bool     `xml:"verify"`
		Amount        int      `xml:"amount"`
		OrderSource   string   `xml:"orderSource"`
		BillToAddress Address  `xml:"billToAddress"`
		Echeck        Echeck   `xml:"echeck"`
	}

	EcheckCredit struct {
		XMLName     xml.Name `xml:"echeckCredit"`
		Id          string   `xml:"id,attr"`
		ReportGroup string   `xml:"reportGroup,attr"`
		CustomerId  string   `xml:"customerId,attr"`
		LitleTxnId  string   `xml:"litleTxnId"`
		Amount      int      `xml:"amount"`
	}

	EcheckVoid struct {
		XMLName     xml.Name `xml:"echeckVoid"`
		Id          string   `xml:"id,attr"`
		ReportGroup string   `xml:"reportGroup,attr"`
		LitleTxnId  string   `xml:"litleTxnId"`
	}

	Sale struct {
		XMLName                  xml.Name                  `xml:"sale"`
		Id                       string                    `xml:"id,attr"`
		ReportGroup              string                    `xml:"reportGroup,attr"`
		CustomerId               string                    `xml:"customerId,attr"`
		OrderId                  string                    `xml:"orderId"`
		Amount                   int                       `xml:"amount"`
		OrderSource              string                    `xml:"orderSource"`
		BillToAddress            Address                   `xml:"billToAddress"`
		Card                     Card                      `xml:"card"`
		CardholderAuthentication *CardholderAuthentication `xml:"cardholderAuthentication"`
		CustomBilling            *CustomBilling            `xml:"customBilling"`
		EnhancedData             *EnhancedData             `xml:"enhancedData"`
	}

	Void struct {
		XMLName     xml.Name `xml:"void"`
		Id          string   `xml:"id,attr"`
		ReportGroup string   `xml:"reportGroup,attr"`
		LitleTxnId  string   `xml:"litleTxnId"`
	}

	Address struct {
		Name         string `xml:"name"`
		AddressLine1 string `xml:"addressLine1"`
		AddressLine2 string `xml:"addressLine2"`
		AddressLine3 string `xml:"addressLine3"`
		City         string `xml:"city"`
		State        string `xml:"state"`
		Zip          string `xml:"zip"`
		Country      string `xml:"country"`
		Email        string `xml:"email"`
		Phone        string `xml:"phone"`
	}

	Echeck struct {
		AccType    string  `xml:"accType"`
		AccNum     string  `xml:"accNum"`
		RoutingNum string  `xml:"routingNum"`
		CheckNum   *string `xml:"checkNum"`
	}

	Card struct {
		Type              string `xml:"type"`
		Number            string `xml:"number"`
		ExpDate           string `xml:"expDate"`
		CardValidationNum string `xml:"cardValidationNum"`
	}

	CardholderAuthentication struct {
		AuthenticationValue         string `xml:"authenticationValue"`
		AuthenticationTransactionId string `xml:"authenticationTransactionId"`
	}

	CustomBilling struct {
		Phone      string `xml:"phone"`
		Descriptor string `xml:"descriptor"`
	}

	EnhancedData struct {
		CustomerReference      string         `xml:"customerReference"`
		SalesTax               int            `xml:"salesTax"`
		TaxExempt              bool           `xml:"taxExempt"`
		DiscountAmount         int            `xml:"discountAmount"`
		ShippingAmount         int            `xml:"shippingAmount"`
		DutyAmount             int            `xml:"dutyAmount"`
		ShipFromPostalCode     string         `xml:"shipFromPostalCode"`
		DestinationPostalCode  string         `xml:"destinationPostalCode"`
		DestinationCountryCode string         `xml:"destinationCountryCode"`
		InvoiceReferenceNumber string         `xml:"invoiceReferenceNumber"`
		OrderDate              string         `xml:"orderDate"`
		DetailTax              DetailTax      `xml:"detailTax"`
		LineItemData           []LineItemData `xml:"lineItemData"`
	}

	DetailTax struct {
		TaxIncludedInTotal bool   `xml:"taxIncludedInTotal"`
		TaxAmount          int    `xml:"taxAmount"`
		TaxRate            string `xml:"taxRate"`
		TaxTypeIdentifier  string `xml:"taxTypeIdentifier"`
		CardAcceptorTaxId  string `xml:"cardAcceptorTaxId"`
	}

	LineItemData struct {
		ItemSequenceNumber   int       `xml:"itemSequenceNumber"`
		ItemDescription      string    `xml:"itemDescription"`
		ProductCode          string    `xml:"productCode"`
		Quantity             int       `xml:"quantity"`
		UnitOfMeasure        string    `xml:"unitOfMeasure"`
		TaxAmount            int       `xml:"taxAmount"`
		LineItemTotal        int       `xml:"lineItemTotal"`
		LineItemTotalWithTax int       `xml:"lineItemTotalWithTax"`
		ItemDiscountAmount   int       `xml:"itemDiscountAmount"`
		CommodityCode        string    `xml:"commodityCode"`
		UnitCost             float64   `xml:"unitCost"`
		DetailTax            DetailTax `xml:"detailTax"`
	}

	AuthorizationResponse struct {
		XMLName              xml.Name        `xml:"authorizationResponse"`
		Id                   string          `xml:"id,attr"`
		ReportGroup          string          `xml:"reportGroup,attr"`
		CustomerId           string          `xml:"customerId,attr"`
		LitleTxnId           string          `xml:"litleTxnId"`
		OrderId              string          `xml:"orderId"`
		Response             string          `xml:"response"`
		ResponseTime         string          `xml:"responseTime"`
		PostDate             string          `xml:"postDate"`
		Message              string          `xml:"message"`
		AuthCode             string          `xml:"authCode"`
		ApprovedAmount       string          `xml:"approvedAmount"`
		NetworkTransactionId string          `xml:"networkTransactionId"`
		FraudResult          *FraudResult    `xml:"fraudResult"`
		AccountUpdater       *AccountUpdater `xml:"accountUpdater"`
	}

	CaptureResponse struct {
		XMLName        xml.Name        `xml:"captureResponse"`
		Id             string          `xml:"id,attr"`
		ReportGroup    string          `xml:"reportGroup,attr"`
		CustomerId     string          `xml:"customerId,attr"`
		LitleTxnId     string          `xml:"litleTxnId"`
		Response       string          `xml:"response"`
		ResponseTime   string          `xml:"responseTime"`
		PostDate       string          `xml:"postDate"`
		Message        string          `xml:"message"`
		AccountUpdater *AccountUpdater `xml:"accountUpdater"`
	}

	CreditResponse struct {
		XMLName      xml.Name `xml:"creditResponse"`
		Id           string   `xml:"id,attr"`
		ReportGroup  string   `xml:"reportGroup,attr"`
		LitleTxnId   string   `xml:"litleTxnId"`
		Response     string   `xml:"response"`
		ResponseTime string   `xml:"responseTime"`
		Message      string   `xml:"message"`
	}

	EcheckCreditResponse struct {
		XMLName        xml.Name        `xml:"echeckCreditResponse"`
		Id             string          `xml:"id,attr"`
		ReportGroup    string          `xml:"reportGroup,attr"`
		CustomerId     string          `xml:"customerId,attr"`
		LitleTxnId     string          `xml:"litleTxnId"`
		Response       string          `xml:"response"`
		ResponseTime   string          `xml:"responseTime"`
		Message        string          `xml:"message"`
		AccountUpdater *AccountUpdater `xml:"accountUpdater"`
	}

	EcheckSaleResponse struct {
		XMLName        xml.Name        `xml:"echeckSalesResponse"`
		Id             string          `xml:"id,attr"`
		ReportGroup    string          `xml:"reportGroup,attr"`
		CustomerId     string          `xml:"customerId,attr"`
		LitleTxnId     string          `xml:"litleTxnId"`
		Response       string          `xml:"response"`
		ResponseTime   string          `xml:"responseTime"`
		Message        string          `xml:"message"`
		PostDate       string          `xml:"postDate"`
		AccountUpdater *AccountUpdater `xml:"accountUpdater"`
	}

	EcheckVoidResponse struct {
		XMLName      xml.Name `xml:"echeckVoidResponse"`
		Id           string   `xml:"id,attr"`
		ReportGroup  string   `xml:"reportGroup,attr"`
		LitleTxnId   string   `xml:"litleTxnId"`
		Response     string   `xml:"response"`
		ResponseTime string   `xml:"responseTime"`
		Message      string   `xml:"message"`
		PostDate     string   `xml:"postDate"`
	}

	SaleResponse struct {
		XMLName              xml.Name        `xml:"saleResponse"`
		Id                   string          `xml:"id,attr"`
		ReportGroup          string          `xml:"reportGroup,attr"`
		CustomerId           string          `xml:"customerId,attr"`
		LitleTxnId           string          `xml:"litleTxnId"`
		Response             string          `xml:"response"`
		OrderId              string          `xml:"orderId"`
		ResponseTime         string          `xml:"responseTime"`
		PostDate             string          `xml:"postDate"`
		Message              string          `xml:"message"`
		AuthCode             string          `xml:"authCode"`
		NetworkTransactionId string          `xml:"networkTransactionId"`
		FraudResult          *FraudResult    `xml:"fraudResult"`
		AccountUpdater       *AccountUpdater `xml:"accountUpdater"`
	}

	VoidResponse struct {
		XMLName      xml.Name `xml:"voidResponse"`
		Id           string   `xml:"id,attr"`
		ReportGroup  string   `xml:"reportGroup,attr"`
		LitleTxnId   string   `xml:"litleTxnId"`
		Response     string   `xml:"response"`
		ResponseTime string   `xml:"responseTime"`
		PostDate     string   `xml:"postDate"`
		Message      string   `xml:"message"`
	}

	FraudResult struct {
		AvsResult            string `xml:"avsResult"`
		CardValidationResult string `xml:"cardValidationResult"`
		AuthenticationResult string `xml:"authenticationResult"`
	}

	AccountUpdater struct {
		OriginalCardInfo Card `xml:"originalCardInfo"`
		NewCardInfo      Card `xml:"newCardInfo"`
	}
)
