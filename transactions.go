package worldpay

import (
	"context"
	"net/http"
)

func (c *Client) Authorization(ctx context.Context, merchantId string, auth *Authorization) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, auth)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Capture(ctx context.Context, merchantId string, capture *Capture) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, capture)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Credit(ctx context.Context, merchantId string, credit *Credit) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, credit)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) EcheckCredit(ctx context.Context, merchantId string, echeckCredit *EcheckCredit) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, echeckCredit)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) EcheckSale(ctx context.Context, merchantId string, echeckSale *EcheckSale) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, echeckSale)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) EcheckVoid(ctx context.Context, merchantId string, echeckVoid *EcheckVoid) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, echeckVoid)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Sale(ctx context.Context, merchantId string, sale *Sale) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, sale)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Void(ctx context.Context, merchantId string, void *Void) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, merchantId, void)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) executeRequest(ctx context.Context, req *http.Request) (*LitleOnlineResponse, error) {
	response := &LitleOnlineResponse{}
	err := c.Send(req, response)
	return response, err
}
