package harperdb

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	endpoint   string
	HttpClient *resty.Client
}

func NewClient(endpoint string, username string, password string) *Client {
	httpClient := resty.
		New().
		SetDisableWarn(true).
		//		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetBasicAuth(username, password)

	return &Client{
		endpoint:   endpoint,
		HttpClient: httpClient,
	}
}

func (c *Client) opRequest(op operation, result interface{}) error {
	e := ErrorResponse{}

	req := c.HttpClient.
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
