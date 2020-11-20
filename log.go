package harperdb

import "time"

const (
	LogOrderAsc  = "asc"
	LogOrderDesc = "desc"

	LogSearchTypeAll       = ""
	LogSearchTypeTimestamp = "timestamp"
	LogSearchTypeUsername  = "username"
	LogSearchTypeHashValue = "hash_value"
)

type LogEntry struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type LogResponse struct {
	File []LogEntry `json:"file"`
}

type TxLogEntry struct {
	Operation  string                   `json:"operation"`
	UserName   string                   `json:"user_name"`
	Timestamp  Timestamp                `json:"timestamp"` // this is a float
	HashValues []interface{}            `json:"hash_values"`
	Records    []map[string]interface{} `json:"records"`
	// It would be possible to pass a target struct
	// to the log read func to enable custom unmarshalling
	// however in a log it might not be possibe to know
	// the exact structure of the data
}

func (c *Client) ReadHarperDBLog(limit, start int, from, until time.Time, order string) (*LogResponse, error) {
	var result LogResponse
	err := c.opRequest(operation{
		Operation: OP_READ_LOG,
		Limit:     limit,
		Start:     0,
		From:      from.Format(DATE_FORMAT),
		Until:     until.Format(DATE_FORMAT),
		Order:     order,
	}, &result)

	return &result, err
}

// ReadTransactionLog requests the transaction log for a table.
// Use LogSearchType* constants to filter the log entries by searchValues,
// which should be an array/slice of searchType.
// Leave searchType empty (LogSearchTypeAll) to get all entries.
func (c Client) ReadTransactionLog(schema, table, searchType string, searchValues interface{}) ([]TxLogEntry, error) {
	var result []TxLogEntry
	err := c.opRequest(operation{
		Operation:    OP_READ_TRANSACTION_LOG,
		Schema:       schema,
		Table:        table,
		SearchType:   searchType,
		SearchValues: searchValues,
	}, &result)
	return result, err
}

func (c Client) DeleteTransactionLogsBefore(schema, table string, timestamp time.Time) error {
	return c.opRequest(operation{
		Operation: OP_DELETE_TRANSACTION_LOG,
		Schema:    schema,
		Table:     table,
		Timestamp: timestamp.UnixNano(),
	}, nil)
}
