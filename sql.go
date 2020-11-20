package harperdb

import "fmt"

// SQLSelect executes SELECT statements and returns a record set.
// You can use format verbs (%s, %d) in the stmt and an pass the arguments at
// the end of the function, like in fmt.Printf
func (c *Client) SQLSelect(v interface{}, stmt string, args ...interface{}) error {
	err := c.opRequest(operation{
		Operation: OP_SQL,
		SQL:       fmt.Sprintf(stmt, args...),
	}, v)
	return err
}

// SQLExec executes UPDATE/INSERT/DELETE statements and returns a struct with
// the affected row hash values.
// You can use format verbs (%s, %d) in the stmt and an pass the arguments at
// the end of the function, like in fmt.Printf
func (c *Client) SQLExec(stmt string, args ...interface{}) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation: OP_SQL,
		SQL:       fmt.Sprintf(stmt, args...),
	}, &result)
	return &result, err
}

// SQLGet is to query a scalar value from the database. This function is
// not part of the official HarperDB API.
// It executes a SQL statement and expects exactly one object with one key.
// I.e. SELECT CURRENT_TIMESTAMP
// Will return the following errors:
// - ErrNoRows
// - ErrTooManyRows
// - ErrNotSingleColumn
func (c *Client) SQLGet(stmt string, args ...interface{}) (interface{}, error) {
	var result []map[string]interface{}

	err := c.opRequest(operation{
		Operation: OP_SQL,
		SQL:       fmt.Sprintf(stmt, args...),
	}, &result)

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrNoRows
	}
	if len(result) > 1 {
		return nil, ErrTooManyRows
	}

	row := result[0]
	for _, val := range row {
		// Return at the first key found...
		return val, nil
	}

	// otherwise, we had either zero or many keys
	return nil, ErrNotSingleColumn
}
