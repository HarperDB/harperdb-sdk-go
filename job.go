package harperdb

import "time"

const (
	JobStatusCompleted  = "COMPLETE"
	JobStatusInProgress = "IN_PROGRESS"
)

type GetJobResponse struct {
	Record
	MessageResponse
	CreatedDateTime        harperTimestamp `json:"created_datetime"`
	EndDateTime            harperTimestamp `json:"end_datetime"`
	StartDateTime          harperTimestamp `json:"start_datetime"`
	ID                     string          `json:"id"`
	JobBody                string          `json:"job_body"` // TODO Not verified
	Status                 string          `json:"status"`
	Type                   string          `json:"type"`
	User                   string          `json:"user"`
	StartDateTimeConverted time.Time       `json:"start_datetime_converted"`
	EndDateTimeConverted   time.Time       `json:"end_datetime_converted"`
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
		return nil, &OperationFailedError{
			err: "expected a job",
		}
	}
	return &resp[0], nil
}
