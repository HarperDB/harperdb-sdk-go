package harperdb

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

type DescribeAllResponse map[string]map[string]DescribeTableResponse

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

func (c *Client) DescribeAll() (*DescribeAllResponse, error) {
	var result DescribeAllResponse
	err := c.opRequest(operation{
		Operation: OP_DESCRIBE_ALL,
	}, &result)
	return &result, err
}

func (c *Client) CreateAttribute(schema, table, attribute string) error {
	return c.opRequest(operation{
		Operation: OP_CREATE_ATTRIBUTE,
		Schema:    schema,
		Table:     table,
		Attribute: attribute,
	}, nil)
}

func (c *Client) DropAttribute(schema, table, attribute string) error {
	return c.opRequest(operation{
		Operation: OP_DROP_ATTRIBUTE,
		Schema:    schema,
		Table:     table,
		Attribute: attribute,
	}, nil)
}
