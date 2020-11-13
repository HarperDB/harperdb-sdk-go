package harperdb

/*
- func (c *Client) CreateSchema(schema string) error
- func (c *Client) DropSchema(schema string) error
- func (c *Client) DescribeSchema(schema string) (SchemaDescription, error)
- func (c *Client) CreateTable(schema, table string, hashAttribute string) (error)
- func (c *Client) DescribeTable(schema, table string) (TableDescription, error)
- func (c *Client) DescribeAll() (DatabaseDescription, error)
- func (c *Client) DropTable(schema, table string) error
- func (c *Client) CreateAttribute(schema, table, attribute string) error
- func (c *Client) DropAttribute(schea, table, attribute string) error
*/

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
