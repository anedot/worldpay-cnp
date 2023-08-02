package worldpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	login      = "username"
	password   = "password"
	apiBase    = "https://www.testvantivcnp.com/sandbox/communicator/online"
	merchantId = "100"
)

func TestGetTransactionXmlPresence(t *testing.T) {
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

	c, _ := NewClient(login, password, apiBase)
	res, _ := c.GetTransactionXml(merchantId, echeckSale)

	assert.NotNil(t, res)
}
