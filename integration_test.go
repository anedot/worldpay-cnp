package worldpay

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorization(t *testing.T) {
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

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.Authorization(context.Background(), auth)
	assert.Equal(t, false, res.HasError())
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

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.Capture(context.Background(), capture)
	assert.Equal(t, false, res.HasError())
}

func TestRefundReversal(t *testing.T) {
	refundReversal := &RefundReversal{
		Id:          "12345",
		CustomerId:  "Customer Id",
		ReportGroup: "Refund Reversals",
		LitleTxnId:  "1234567890123456789",
		Card: Card{
			Type:              "GC",
			Number:            "1234102000003558",
			CardValidationNum: "888",
			ExpDate:           "1210",
		},
		OriginalRefCode:        "123456",
		OriginalAmount:         1900,
		OriginalTxnTime:        "2017-03-21T10:02:46",
		OriginalSystemTraceId:  "678901",
		OriginalSequenceNumber: "123456",
	}

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.RefundReversal(context.Background(), refundReversal)
	assert.Equal(t, false, res.HasError())
}

func TestSale(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
				Number:            "4005550000081019",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		}

		c, _ := NewClient(login, password, apiBase, merchantId)
		res, _ := c.Sale(context.Background(), sale)
		assert.Equal(t, false, res.HasError())
	})

	t.Run("invalid card type", func(t *testing.T) {
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

		c, _ := NewClient(login, password, apiBase, merchantId)
		res, _ := c.Sale(context.Background(), sale)
		assert.Equal(t, true, res.HasError())
		assert.Equal(
			t,
			"Error validating xml data against the schema: cvc-enumeration-valid: Value 'FOO' is not facet-valid with respect to enumeration '[MC, VI, AX, DC, DI, PP, JC, BL, EC, GC, ]'. It must be a value from the enumeration.",
			res.Error(),
		)

	})

	t.Run("invalid card validation num", func(t *testing.T) {
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
				Number:            "4005550000081019",
				ExpDate:           "1210",
				CardValidationNum: "55511111111111111",
			},
		}

		c, _ := NewClient(login, password, apiBase, merchantId)
		res, _ := c.Sale(context.Background(), sale)
		assert.Equal(t, true, res.HasError())
		assert.Equal(
			t,
			"Error validating xml data against the schema: cvc-maxLength-valid: Value '55511111111111111' with length = '17' is not facet-valid with respect to maxLength '4' for type 'cvNumType'.",
			res.Error(),
		)

	})
}

func TestVoid(t *testing.T) {
	void := &Void{
		Id:          "12345",
		ReportGroup: "report group",
		LitleTxnId:  "1234567890123456789",
	}

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.Void(context.Background(), void)
	assert.Equal(t, false, res.HasError())
}

func TestEcheckSale(t *testing.T) {
	echeckSale := &EcheckSale{
		Id:          "1",
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

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.EcheckSale(context.Background(), echeckSale)
	assert.Equal(t, false, res.HasError())
}

func TestEcheckCredit(t *testing.T) {
	echeckCredit := &EcheckCredit{
		Id:          "credit1",
		ReportGroup: "new53",
		CustomerId:  "53",
		LitleTxnId:  "4455667788",
		Amount:      1000,
	}

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.EcheckCredit(context.Background(), echeckCredit)
	assert.Equal(t, false, res.HasError())
}

func TestEcheckVoid(t *testing.T) {
	echeckVoid := &EcheckVoid{
		Id:          "101",
		ReportGroup: "001601",
		LitleTxnId:  "345454444",
	}

	c, _ := NewClient(login, password, apiBase, merchantId)
	res, _ := c.EcheckVoid(context.Background(), echeckVoid)
	assert.Equal(t, false, res.HasError())
}
