package worldpay

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSale(t *testing.T) {
	t.Run("Code: 000", func(t *testing.T) {
		sale := &Sale{
			Id:          "1",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4005550000081000",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Sale(context.Background(), merchantId, sale)
		assert.Equal(t, false, res.HasError())
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "1", res.SaleResponse.Id)
		assert.Equal(t, "ABC Division", res.SaleResponse.ReportGroup)
		assert.Equal(t, "", res.SaleResponse.CustomerId)
		assert.Equal(t, "000", res.SaleResponse.Response)
		assert.Equal(t, "5234234", res.SaleResponse.OrderId)
		assert.Equal(t, "Approved", res.SaleResponse.Message)
		assert.Equal(t, nil, res.SaleResponse.FraudResult)
	})

	t.Run("with AVS", func(t *testing.T) {
		sale := &Sale{
			Id:          "1",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4200410886320101",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Sale(context.Background(), merchantId, sale)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "1", res.SaleResponse.Id)
		assert.Equal(t, "ABC Division", res.SaleResponse.ReportGroup)
		assert.Equal(t, "", res.SaleResponse.CustomerId)
		assert.Equal(t, "101", res.SaleResponse.Response)
		assert.Equal(t, "5234234", res.SaleResponse.OrderId)
		assert.Equal(t, "Issuer Unavailable", res.SaleResponse.Message)
		assert.Equal(t, "10", res.SaleResponse.FraudResult.AvsResult)
		assert.Equal(t, "", res.SaleResponse.FraudResult.CardValidationResult)
		assert.Equal(t, "", res.SaleResponse.FraudResult.AuthenticationResult)
	})

	t.Run("with Card Security Code Validation", func(t *testing.T) {
		sale := &Sale{
			Id:          "1",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4100521234567000",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Sale(context.Background(), merchantId, sale)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "1", res.SaleResponse.Id)
		assert.Equal(t, "ABC Division", res.SaleResponse.ReportGroup)
		assert.Equal(t, "", res.SaleResponse.CustomerId)
		assert.Equal(t, "000", res.SaleResponse.Response)
		assert.Equal(t, "5234234", res.SaleResponse.OrderId)
		assert.Equal(t, "Approved", res.SaleResponse.Message)
		assert.Equal(t, "", res.SaleResponse.FraudResult.AvsResult)
		assert.Equal(t, "P", res.SaleResponse.FraudResult.CardValidationResult)
		assert.Equal(t, "", res.SaleResponse.FraudResult.AuthenticationResult)
	})

	t.Run("with Account Updater", func(t *testing.T) {
		sale := &Sale{
			Id:          "1",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4100117890123000",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Sale(context.Background(), merchantId, sale)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "1", res.SaleResponse.Id)
		assert.Equal(t, "VI", res.SaleResponse.AccountUpdater.OriginalCardInfo.Type)
		assert.Equal(t, "4100117890123000", res.SaleResponse.AccountUpdater.OriginalCardInfo.Number)
		assert.Equal(t, "1210", res.SaleResponse.AccountUpdater.OriginalCardInfo.ExpDate)
		assert.Equal(t, "VI", res.SaleResponse.AccountUpdater.NewCardInfo.Type)
		assert.Equal(t, "1210", res.SaleResponse.AccountUpdater.NewCardInfo.ExpDate)
	})

	t.Run("with validation error", func(t *testing.T) {
		sale := &Sale{
			Id:          "1",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "FOO",
				Number:            "4005550000081019",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Sale(context.Background(), merchantId, sale)
		assert.Equal(t, true, res.HasError())
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "1", res.Response)
		assert.Equal(t,
			"Error validating xml data against the schema: cvc-enumeration-valid: Value 'FOO' is not facet-valid with respect to enumeration '[MC, VI, AX, DC, DI, PP, JC, BL, EC, GC, ]'. It must be a value from the enumeration.",
			res.Message,
		)
		assert.Equal(t, nil, res.SaleResponse)
	})
}

func TestAuthorization(t *testing.T) {
	t.Run("Approved", func(t *testing.T) {
		auth := &Authorization{
			Id:          "834262",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "65347567",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				City:         "Boston",
				Country:      "USA",
				State:        "MA",
				Zip:          "12345",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4000000000000002",
				ExpDate:           "1209",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Authorization(context.Background(), merchantId, auth)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "834262", res.AuthorizationResponse.Id)
		assert.Equal(t, "000", res.AuthorizationResponse.Response)
		assert.Equal(t, "Approved", res.AuthorizationResponse.Message)
		assert.Equal(t, nil, res.AuthorizationResponse.AccountUpdater)
	})

	t.Run("with AVS", func(t *testing.T) {
		authorization := &Authorization{
			Id:          "1",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			OrderId:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4200410886320101",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Authorization(context.Background(), merchantId, authorization)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "1", res.AuthorizationResponse.Id)
		assert.Equal(t, "ABC Division", res.AuthorizationResponse.ReportGroup)
		assert.Equal(t, "", res.AuthorizationResponse.CustomerId)
		assert.Equal(t, "101", res.AuthorizationResponse.Response)
		assert.Equal(t, "5234234", res.AuthorizationResponse.OrderId)
		assert.Equal(t, "Issuer Unavailable", res.AuthorizationResponse.Message)
		assert.Equal(t, "10", res.AuthorizationResponse.FraudResult.AvsResult)
		assert.Equal(t, "", res.AuthorizationResponse.FraudResult.CardValidationResult)
		assert.Equal(t, "", res.AuthorizationResponse.FraudResult.AuthenticationResult)
	})
}

func TestCapture(t *testing.T) {
	capture := &Capture{
		Id:          "834262",
		ReportGroup: "ABC Division",
		CustomerId:  "038945",
		Partial:     false,
		LitleTxnId:  "13254123434",
		Amount:      5000,
	}

	c, _ := NewClient(login, password, apiBase)
	res, _ := c.Capture(context.Background(), merchantId, capture)
	assert.Equal(t, "11.4", res.Version)
	assert.Equal(t, "0", res.Response)
	assert.Equal(t, "Valid Format", res.Message)
	assert.Equal(t, "834262", res.CaptureResponse.Id)
	assert.Equal(t, "000", res.CaptureResponse.Response)
	assert.Equal(t, "Approved", res.CaptureResponse.Message)
	assert.Equal(t, nil, res.CaptureResponse.AccountUpdater)
}

func TestCredit(t *testing.T) {
	t.Run("with amount given", func(t *testing.T) {
		amount := 5000

		credit := &Credit{
			Id:          "834262",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			LitleTxnId:  "13254123434",
			Amount:      &amount,
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Credit(context.Background(), merchantId, credit)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "834262", res.CreditResponse.Id)
		assert.Equal(t, "000", res.CreditResponse.Response)
		assert.Equal(t, "Approved", res.CreditResponse.Message)
	})

	t.Run("without amount given", func(t *testing.T) {
		credit := &Credit{
			Id:          "834262",
			ReportGroup: "ABC Division",
			CustomerId:  "038945",
			LitleTxnId:  "13254123434",
		}

		c, _ := NewClient(login, password, apiBase)
		res, _ := c.Credit(context.Background(), merchantId, credit)
		assert.Equal(t, "11.4", res.Version)
		assert.Equal(t, "0", res.Response)
		assert.Equal(t, "Valid Format", res.Message)
		assert.Equal(t, "834262", res.CreditResponse.Id)
		assert.Equal(t, "000", res.CreditResponse.Response)
		assert.Equal(t, "Approved", res.CreditResponse.Message)
	})
}

func TestVoid(t *testing.T) {
	void := &Void{
		Id:          "834262",
		ReportGroup: "report group",
		LitleTxnId:  "1234567890123456789",
	}

	c, _ := NewClient(login, password, apiBase)
	res, _ := c.Void(context.Background(), merchantId, void)
	assert.Equal(t, "11.4", res.Version)
	assert.Equal(t, "0", res.Response)
	assert.Equal(t, "Valid Format", res.Message)
	assert.Equal(t, "834262", res.VoidResponse.Id)
	assert.Equal(t, "000", res.VoidResponse.Response)
	assert.Equal(t, "Approved", res.VoidResponse.Message)
}

func TestEcheckSale(t *testing.T) {
	echeckSale := &EcheckSale{
		Id:          "834262",
		ReportGroup: "ABC Division",
		CustomerId:  "038945",
		OrderId:     "5234234",
		Amount:      40000,
		Verify:      false,
		OrderSource: "3dsAuthenticated",
		BillToAddress: Address{
			Name:         "John Smith",
			AddressLine1: "100 Main St",
			AddressLine2: "100 Main St",
			AddressLine3: "100 Main St",
			City:         "Boston",
			State:        "MA",
			Zip:          "12345",
			Country:      "US",
			Email:        "jsmith@someaddress.com",
			Phone:        "555-123-4567",
		},
		Echeck: Echeck{
			AccType:    "Checking",
			AccNum:     "5186005800001012",
			RoutingNum: "000010101",
		},
	}

	c, _ := NewClient(login, password, apiBase)
	res, _ := c.EcheckSale(context.Background(), merchantId, echeckSale)
	assert.Equal(t, "11.4", res.Version)
	assert.Equal(t, "0", res.Response)
	assert.Equal(t, "Valid Format", res.Message)
	assert.Equal(t, "834262", res.EcheckSaleResponse.Id)
	assert.Equal(t, "000", res.EcheckSaleResponse.Response)
	assert.Equal(t, "Approved", res.EcheckSaleResponse.Message)
}

func TestEcheckCredit(t *testing.T) {
	echeckCredit := &EcheckCredit{
		Id:          "834262",
		ReportGroup: "new53",
		CustomerId:  "53",
		LitleTxnId:  "4455667788",
		Amount:      1000,
	}

	c, _ := NewClient(login, password, apiBase)
	res, _ := c.EcheckCredit(context.Background(), merchantId, echeckCredit)
	assert.Equal(t, "11.4", res.Version)
	assert.Equal(t, "0", res.Response)
	assert.Equal(t, "Valid Format", res.Message)
	assert.Equal(t, "834262", res.EcheckCreditResponse.Id)
	assert.Equal(t, "000", res.EcheckCreditResponse.Response)
	assert.Equal(t, "Approved", res.EcheckCreditResponse.Message)
}

func TestEcheckVoid(t *testing.T) {
	echeckVoid := &EcheckVoid{
		Id:          "834262",
		ReportGroup: "001601",
		LitleTxnId:  "345454444",
	}

	c, _ := NewClient(login, password, apiBase)
	res, _ := c.EcheckVoid(context.Background(), merchantId, echeckVoid)
	assert.Equal(t, "11.4", res.Version)
	assert.Equal(t, "0", res.Response)
	assert.Equal(t, "Valid Format", res.Message)
	assert.Equal(t, "834262", res.EcheckVoidResponse.Id)
	assert.Equal(t, "000", res.EcheckVoidResponse.Response)
	assert.Equal(t, "Approved", res.EcheckVoidResponse.Message)
}
