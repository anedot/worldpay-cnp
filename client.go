package worldpay

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

const (
	version      = "11.4"
	xmlNamespace = "http://www.litle.com/schema"
)

func NewClient(login, password, apiBase, merchantId string) (*Client, error) {
	if login == "" || password == "" || apiBase == "" || merchantId == "" {
		return nil, errors.New("Missing required credentials")
	}

	return &Client{
		Client:     &http.Client{},
		Login:      login,
		Password:   password,
		ApiBase:    apiBase,
		MerchantId: merchantId,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	// default headers
	req.Header.Set("Content-Type", "text/xml")

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return xml.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

// log will dump request and response to the log file
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}

func (c *Client) GetTransactionXml(payload interface{}) ([]byte, error) {
	request := LitleOnlineRequest{
		Version:      version,
		XmlNamespace: xmlNamespace,
		MerchantId:   c.MerchantId,
		Authentication: Authentication{
			User:     c.Login,
			Password: c.Password,
		},
	}

	// Use type assertion to check the type of the payload
	switch p := payload.(type) {
	case *Authorization:
		request.Authorization = p
	case *Capture:
		request.Capture = p
	case *EcheckCredit:
		request.EcheckCredit = p
	case *EcheckSale:
		request.EcheckSale = p
	case *EcheckVoid:
		request.EcheckVoid = p
	case *Sale:
		request.Sale = p
	case *RefundReversal:
		request.RefundReversal = p
	case *Void:
		request.Void = p
	}

	return xml.MarshalIndent(request, "", "  ")
}

func (c *Client) NewRequest(ctx context.Context, payload interface{}) (*http.Request, error) {
	xmlData, _ := c.GetTransactionXml(payload)

	return http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.ApiBase,
		bytes.NewReader(xmlData),
	)
}
