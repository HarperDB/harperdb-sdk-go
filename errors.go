package harperdb

type OperationFailedError struct {
	err string
}

func (e *OperationFailedError) Error() string {
	return e.err
}

type AlreadyExistsError struct {
	OperationFailedError
}

type DoesNotExistsError struct {
	OperationFailedError
}
