package harperdb

/*
- func (c *Client) CreateTable(schema, table string, hashAttribute string) (error)
- func (c *Client) DropTable(schema, table string) error
- func (c *Client) DescribeTable(schema, table string) (TableDescription, error)
- func (c *Client) DescribeAll() (DatabaseDescription, error)
- func (c *Client) CreateAttribute(schema, table, attribute string) error
- func (c *Client) DropAttribute(schea, table, attribute string) error
*/

func (c *Client) CreateTable(schema, table, hashAttribute string) error {
	return c.opRequest(operation{
		Operation:     OP_CREATE_TABLE,
		Schema:        schema,
		Table:         table,
		HashAttribute: hashAttribute,
	}, nil)
}

func (c *Client) DropTable(schema, table, hashAttribute string) error {
	return c.opRequest(operation{
		Operation: OP_DROP_TABLE,
		Schema:    schema,
		Table:     table,
	}, nil)
}
