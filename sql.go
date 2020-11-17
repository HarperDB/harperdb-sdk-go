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
