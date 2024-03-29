# worldpay-cnp

> SCHEMA VERSION 11.4 - Go interface to the [WorldPay cnpAPI](http://support.worldpay.com/support/CNP-API/content/introduction.htm)

## API

```go
func NewClient(login, password, url, merchant_id string) Client
```

Client
```go
func (c *Client) Authorization(c Context, authorization *Authorization) LitleOnlineResponse
func (c *Client) Capture(c Context, capture *Capture) LitleOnlineResponse
func (c *Client) Credit(c Context, credit *Credit) LitleOnlineResponse
func (c *Client) EcheckCredit(c Context, echeckCredit *EcheckCredit) LitleOnlineResponse
func (c *Client) EcheckSale(c Context, echeckSale *EcheckSale) LitleOnlineResponse
func (c *Client) EcheckVoid(c Context, echeckVoid *EcheckVoid) LitleOnlineResponse
func (c *Client) Sale(c Context, sale *Sale) LitleOnlineResponse
func (c *Client) Void(c Context, void *Void) LitleOnlineResponse
```

## Usage
```go
import (
    "context"
    "os"

    "github.com/anedot/worldpay-cnp"
)

func main() {
    client, _ := worldpay.NewClient(
        os.Getenv("WORLDPAY_LOGIN"),
        os.Getenv("WORLDPAY_PASSWORD"),
        os.Getenv("WORLDPAY_URL"),
        os.Getenv("WORLDPAY_MERCHANT_ID"),
    )

    ctx := context.Background()
    void := &worldpay.Void{
        Id:          "12345",
        ReportGroup: "report group",
        LitleTxnId:  "1234567890123456789",
    }

    client.Void(ctx, void)
}
```

## Online Transactions

The wrapper currently only supports "online" transactions.
### Authorization

```go
func Authorization(c Context, authorization *Authorization) LitleOnlineResponse
```

```go
&worldpay.Authorization{
    Id:          "12345",
    ReportGroup: "ABC Division",
    CustomerId:  "038945",
    OrderId:     "65347567",
    Amount:      40000,
    OrderSource: "3dsAuthenticated",
    BillToAddress: worldpay.Address{
        Name:         "John Smith",
        AddressLine1: "100 Main St",
        City:         "Boston",
        Country:      "USA",
        State:        "MA",
        Zip:          "12345",
        Email:        "jsmith@someaddress.com",
        Phone:        "555-123-4567",
    },
    Card: worldpay.Card{
        Type:              "VI",
        Number:            "4000000000000002",
        ExpDate:           "1209",
        CardValidationNum: "555",
    },
}
```

### Capture
```go
func Capture(c Context, capture *Capture) LitleOnlineResponse
```

```go
&worldpay.Capture{
    Id:          "12345",
    ReportGroup: "ABC Division",
    CustomerId:  "038945",
    Partial:     false,
    LitleTxnId:  "13254123434",
    Amount:      5000,
}
```

### Credit
```go
func Credit(c Context, credit *Credit) LitleOnlineResponse
```

* With Amount

```go
amount := 5000

&worldpay.Credit{
    Id:          "12345",
    ReportGroup: "ABC Division",
    CustomerId:  "038945",
    Partial:     false,
    LitleTxnId:  "13254123434",
    Amount:      &amount,
}
```

* Without Amount

```go
&worldpay.Credit{
    Id:          "12345",
    ReportGroup: "ABC Division",
    CustomerId:  "038945",
    Partial:     false,
    LitleTxnId:  "13254123434",
}
```

### ECheck Credit
```go
func EcheckCredit(c Context, echeckCredit *EcheckCredit) LitleOnlineResponse
```

```go
&worldpay.EcheckCredit{
    ID:          "credit1",
    ReportGroup: "new53",
    CustomerID:  "53",
    LitleTxnID:  "4455667788",
    Amount:      1000,
}
```

### ECheck Void
```go
func EcheckVoid(c Context, echeckVoid *EcheckVoid) LitleOnlineResponse
```

```go
&worldpay.EcheckVoid{
    Id:          "101",
    ReportGroup: "001601",
    LitleTxnId:  "345454444",
}
```

### ECheck Sale
```go
func EcheckSale(c Context, echeckSale *EcheckSale) LitleOnlineResponse
```

```go
&worldpay.EcheckSale{
    Id:          "1",
    ReportGroup: "ABC Division",
    CustomerId:  "038945",
    OrderId:     "5234234",
    Amount:      40000,
    Verify:      false,
    OrderSource: "3dsAuthenticated",
    BillToAddress: worldpay.Address{
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
    Echeck: worldpay.Echeck{
        AccType:    "Checking",
        AccNum:     "5186005800001012",
        RoutingNum: "000010101",
    },
}
```

### Sale
```go
func Sale(c Context, sale *Sale) LitleOnlineResponse
```

```go
&worldpay.Sale{
    Id:          "1",
    ReportGroup: "ABC Division",
    CustomerId:  "038945",
    OrderId:     "5234234",
    Amount:      40000,
    OrderSource: "3dsAuthenticated",
    BillToAddress: &worldpay.Address{
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
    Card: &Card{
        Type:              "VI",
        Number:            "4005550000081019",
        ExpDate:           "1210",
        CardValidationNum: "555",
    },
}
```

### Void
```go
func Void(c Context, void *Void) LitleOnlineResponse
```

```go
&worldpay.Void{
    Id:          "12345",
    ReportGroup: "report group",
    LitleTxnId:  "1234567890123456789",
}
```

## Dev
### Run tests
```bash
go test ./...
```
