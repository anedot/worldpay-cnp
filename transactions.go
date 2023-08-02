package worldpay

import (
	"context"
	"net/http"
)

func (c *Client) Authorization(ctx context.Context, auth *Authorization) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, auth)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Capture(ctx context.Context, capture *Capture) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, capture)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Credit(ctx context.Context, credit *Credit) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, credit)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) EcheckCredit(ctx context.Context, echeckCredit *EcheckCredit) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, echeckCredit)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) EcheckSale(ctx context.Context, echeckSale *EcheckSale) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, echeckSale)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) EcheckVoid(ctx context.Context, echeckVoid *EcheckVoid) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, echeckVoid)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Sale(ctx context.Context, sale *Sale) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, sale)
	if err != nil {
		return nil, err
	}
	return c.executeRequest(ctx, req)
}

func (c *Client) Void(ctx context.Context, void *Void) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, void)
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
