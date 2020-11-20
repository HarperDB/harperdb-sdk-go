package harperdb

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	CSV_ACTION_INSERT = "insert"
	CSV_ACTION_UPDATE = "update"
)

type JobResponse struct {
	JobID string
}

// CSVDataLoad takes a Reader and executes the CSV Load Data operation
// if "update" is true, it will not insert but update existing records
// If successfull, returns the Job ID
func (c *Client) CSVDataLoad(schema, table string, update bool, data io.Reader) (string, error) {
	return c.csvDataLoad(schema, table, actionByBool(update), data)
}

func (c *Client) csvDataLoad(schema, table, action string, data io.Reader) (string, error) {
	resp := MessageResponse{}
	buff, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}

	err = c.opRequest(operation{
		Operation: OP_CSV_DATA_LOAD,
		Action:    CSV_ACTION_INSERT,
		Schema:    schema,
		Table:     table,
		Data:      string(buff),
	}, &resp)
	if err != nil {
		return "", err
	}

	if jobID, ok := extractJobIDFromMessage(resp.Message); ok {
		return jobID, nil
	}

	return "", fmt.Errorf("did not get a job ID from harper instance: %w", ErrJobStatusUnknown)
}

// CSVFileLoad takes a path of a file which must exist on the server
// and executes the CSV Load Data operation
// if "update" is true, it will not insert but update existing records
// If successfull, returns the Job ID
func (c *Client) CSVFileLoad(schema, table string, update bool, filePath string) (string, error) {
	resp := MessageResponse{}

	err := c.opRequest(operation{
		Operation: OP_CSV_DATA_LOAD,
		Action:    CSV_ACTION_INSERT,
		Schema:    schema,
		Table:     table,
		FilePath:  filePath,
	}, &resp)
	if err != nil {
		return "", err
	}

	if jobID, ok := extractJobIDFromMessage(resp.Message); ok {
		return jobID, nil
	}

	return "", fmt.Errorf("did not get a job ID from harper instance: %w", ErrJobStatusUnknown)
}

// CSVURLLoad takes a public URL
// and executes the CSV Load Data operation
// if "update" is true, it will not insert but update existing records
// If successfull, returns the Job ID
func (c *Client) CSVURLLoad(schema, table string, update bool, csvURL string) (string, error) {
	resp := MessageResponse{}

	err := c.opRequest(operation{
		Operation: OP_CSV_DATA_LOAD,
		Action:    CSV_ACTION_INSERT,
		Schema:    schema,
		Table:     table,
		CSVURL:    csvURL,
	}, &resp)
	if err != nil {
		return "", err
	}

	if jobID, ok := extractJobIDFromMessage(resp.Message); ok {
		return jobID, nil
	}

	return "", fmt.Errorf("did not get a job ID from harper instance: %w", ErrJobStatusUnknown)
}

// extractJobIDFromMessage Askes HarperDB team to return the Job ID
// structured in the JSON response, for now we parse
func extractJobIDFromMessage(message string) (string, bool) {
	jobID := strings.Replace(message, "Starting job with id ", "", 1)
	if len(jobID) != 36 {
		return "", false
	}
	return jobID, true
}

func actionByBool(update bool) string {
	if update {
		return CSV_ACTION_UPDATE
	}
	return CSV_ACTION_INSERT
}
