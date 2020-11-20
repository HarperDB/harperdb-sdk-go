package harperdb

// Low-level operations

// CreateSchema creates a new schema.
// Returns "AlreadyExistsError" if schema already existed.
func (c *Client) CreateSchema(schema string) error {
	return c.opRequest(operation{
		Operation: OP_CREATE_SCHEMA,
		Schema:    schema,
	}, nil)
}

// DropSchema drops a schema.
// Returns "DoesNotExistError" if schema did not exist.
func (c *Client) DropSchema(schema string) error {
	return c.opRequest(operation{
		Operation: OP_DROP_SCHEMA,
		Schema:    schema,
	}, nil)
}

type DescribeSchemaResponse struct {
}

// DescribeSchema returns metadata about a schema.
/*
func (c *Client) DescribeSchema(schema string) (DescribeSchemaResponse, error) {

}
*/
