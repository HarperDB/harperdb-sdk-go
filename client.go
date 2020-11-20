package harperdb

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	endpoint   string
	httpClient *resty.Client
}

func NewClient(endpoint string, username string, password string) *Client {
	httpClient := resty.
		New().
		SetDisableWarn(true).
		SetBasicAuth(username, password)

	return &Client{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
}

func (c *Client) opRequest(op operation, result interface{}) error {
	e := ErrorResponse{}

	req := c.httpClient.
		NewRequest().
		SetBody(op).
		SetError(&e)

	if result != nil {
		req.SetResult(result)
	}

	resp, err := req.Post(c.endpoint)
	if err != nil {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    err.Error()}
	}
	if resp.StatusCode() > 399 {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    e.Error}
	}

	return nil
}
