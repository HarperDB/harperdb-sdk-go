package harperdb

import (
	"fmt"
	"time"
)

const (
	JobStatusCompleted  = "COMPLETE"
	JobStatusInProgress = "IN_PROGRESS"

	DATE_FORMAT = "2006-01-02"
)

type GetJobResponse struct {
	Record
	MessageResponse
	CreatedDateTime        Timestamp `json:"created_datetime"`
	EndDateTime            Timestamp `json:"end_datetime"`
	StartDateTime          Timestamp `json:"start_datetime"`
	ID                     string    `json:"id"`
	JobBody                string    `json:"job_body"` // TODO Not verified
	Status                 string    `json:"status"`
	Type                   string    `json:"type"`
	User                   string    `json:"user"`
	StartDateTimeConverted time.Time `json:"start_datetime_converted"`
	EndDateTimeConverted   time.Time `json:"end_datetime_converted"`
}

func (c *Client) GetJob(jobID string) (*GetJobResponse, error) {
	resp := []GetJobResponse{} // (sic) returns an array, not a single job

	if err := c.opRequest(operation{
		Operation: OP_GET_JOB,
		ID:        jobID,
	}, &resp); err != nil {
		return nil, err
	}
	if len(resp) != 1 {
		return nil, fmt.Errorf("get job: %w", ErrJobNotFound)
	}
	return &resp[0], nil
}

func (c *Client) SearchJobsByStartDate(fromDate, toDate time.Time) ([]GetJobResponse, error) {
	resp := []GetJobResponse{}

	if err := c.opRequest(operation{
		Operation: OP_SEARCH_JOBS,
		FromDate:  fromDate.Format(DATE_FORMAT),
		ToDate:    toDate.Format(DATE_FORMAT),
	}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
