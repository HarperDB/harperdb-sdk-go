package harperdb

/*
- func (c *Client) CreateTable(schema, table string, hashAttribute string) (error)
- func (c *Client) DropTable(schema, table string) error
- func (c *Client) DescribeTable(schema, table string) (TableDescription, error)
- func (c *Client) DescribeAll() (DatabaseDescription, error)
- func (c *Client) CreateAttribute(schema, table, attribute string) error
- func (c *Client) DropAttribute(schea, table, attribute string) error
*/

type DescribeTableResponse struct {
	Record
	HashAttribute string           `json:"hash_attribute"`
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	Residence     string           `json:"residence"` // TODO Not verified
	Schema        string           `json:"schema"`
	Attributes    []TableAttribute `json:"attributes"`
	RecordCount   int              `json:"record_count"`
}

type TableAttribute struct {
	Attribute string `json:"attribute"`
}

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

func (c *Client) DescribeTable(schema, table string) (*DescribeTableResponse, error) {
	resp := DescribeTableResponse{}
	err := c.opRequest(operation{
		Operation: OP_DESCRIBE_TABLE,
		Table:     table,
		Schema:    schema,
	}, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
