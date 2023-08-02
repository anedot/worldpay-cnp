package worldpay

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

const (
	version      = "11.4"
	xmlNamespace = "http://www.litle.com/schema"
)

func NewClient(login, password, apiBase string) (*Client, error) {
	if login == "" || password == "" || apiBase == "" {
		return nil, errors.New("Missing required credentials")
	}

	return &Client{
		Client:   &http.Client{},
		Login:    login,
		Password: password,
		ApiBase:  apiBase,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	// default headers
	req.Header.Set("Content-Type", "text/xml")

	// Read the request body
	reqBody, _ := ioutil.ReadAll(req.Body)
	// Create a new buffer with the request body content
	bodyBuffer := bytes.NewBuffer(reqBody)
	// Reset the request body for the subsequent request
	req.Body = ioutil.NopCloser(bodyBuffer)

	resp, err = c.Client.Do(req)
	c.log(req, reqBody, resp)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return xml.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

func (c *Client) log(r *http.Request, reqBody []byte, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s \n\n%s", r.Method, r.URL.String(), string(reqBody))
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\n\n\nResponse: %s\n", reqDump, string(respDump))))
	}
}

func (c *Client) GetTransactionXml(merchantId string, payload interface{}) ([]byte, error) {
	request := LitleOnlineRequest{
		Version:      version,
		XmlNamespace: xmlNamespace,
		MerchantId:   merchantId,
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
	case *Credit:
		request.Credit = p
	case *EcheckCredit:
		request.EcheckCredit = p
	case *EcheckSale:
		request.EcheckSale = p
	case *EcheckVoid:
		request.EcheckVoid = p
	case *Sale:
		request.Sale = p
	case *Void:
		request.Void = p
	}

	return xml.MarshalIndent(request, "", "  ")
}

func (c *Client) NewRequest(ctx context.Context, merchantId string, payload interface{}) (*http.Request, error) {
	xmlData, _ := c.GetTransactionXml(merchantId, payload)

	return http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.ApiBase,
		bytes.NewReader(xmlData),
	)
}
