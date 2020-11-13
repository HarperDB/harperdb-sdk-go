package harperdb

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	endpoint   string
	httpClient *resty.Client
}

type operation struct {
	Operation string        `json:"operation"`
	Schema    string        `json:"schema,omitempty"`
	Table     string        `json:"table,omitempty"`
	Records   []interface{} `json:"records,omitempty"`
}

func NewClient(endpoint string, username string, password string) *Client {
	httpClient := resty.
		New().
		SetBasicAuth(username, password)

	return &Client{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
}

func (c *Client) opRequest(body interface{}, result interface{}) (*resty.Response, error) {
	return c.httpClient.NewRequest().SetBody(body).SetResult(result).Post(c.endpoint)
}
