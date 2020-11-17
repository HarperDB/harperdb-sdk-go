package harperdb

import (
	"strings"

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
		return &OperationFailedError{err: err.Error()}
	}

	if resp.StatusCode() > 399 {
		if isAlreadyExistsError(e.Error) {
			return &AlreadyExistsError{OperationFailedError: OperationFailedError{err: e.Error}}
		} else if isDoesNotExistsError(e.Error) {
			return &DoesNotExistsError{OperationFailedError: OperationFailedError{err: e.Error}}
		}
		return &OperationFailedError{err: e.Error}
	}

	return nil
}

func isAlreadyExistsError(msg string) bool {
	return strings.Contains(msg, "already exists")
}

func isDoesNotExistsError(msg string) bool {
	return strings.Contains(msg, "does not exist")
}
