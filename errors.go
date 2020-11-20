package harperdb

import (
	"errors"
	"fmt"
	"strings"
)

type OperationError struct {
	StatusCode int
	Message    string
}

func (e *OperationError) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.StatusCode)
}

func (e *OperationError) IsAlreadyExistsError() bool {
	return e.StatusCode > 399 && strings.Contains(e.Message, "already exists")
}

func (e *OperationError) IsDoesNotExistError() bool {
	return e.StatusCode > 399 && strings.Contains(e.Message, "does not exist")
}

func (e *OperationError) IsNotAuthorizedError() bool {
	return e.StatusCode == 403
}

var (
	ErrJobStatusUnknown = errors.New("unknown job status")
	ErrJobNotFound      = errors.New("job not found")
	ErrNoRows           = errors.New("did not return any rows")
	ErrTooManyRows      = errors.New("did return more than one row")
	ErrNotSingleColumn  = errors.New("expected a single column return")
)
