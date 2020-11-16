package harperdb

/*

- func (c *Client) Insert(schema, table string, records []Record) error
- func (c *Client) Update(schema, table string, records []Record) error
- func (c *Client) Delete(schema, table string, hashValues []HashValue) error
- func (c *Client) SearchByHash(schema, table string, hashValues []HashValue, getAttributes []string) ([]Record, error)
- func (c *Client) SearchByValue(schema, table, searchAttribute string, getAttributes []string) ([]Record, error)


*/

// Insert inserts one or more JSON objects into a table
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Insert(schema, table string, records []Record) error {
	if len(records) == 0 {
		// there is nothing to do
		return nil
	}

	return c.opRequest(operation{
		Operation: OP_INSERT,
		Schema:    schema,
		Table:     table,
		Records:   records,
	}, nil)
}

// Update updates one or more JSON objects in a table.
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Update(schema, table string, records []Record) error {
	if len(records) == 0 {
		// there is nothing to do
		return nil
	}

	return c.opRequest(operation{
		Operation: OP_UPDATE,
		Schema:    schema,
		Table:     table,
		Records:   records,
	}, nil)
}

// Delete delete one or more JSON objects from a table.
// hashValues must be an array of slice
func (c *Client) Delete(schema, table string, hashValues AttributeList) error {
	return c.opRequest(operation{
		Operation:  OP_DELETE,
		Schema:     schema,
		Table:      table,
		HashValues: hashValues,
	}, nil)
}

// SearchByHash fetches records based on the table's hash field
// (i.e. primary key).
func (c *Client) SearchByHash(schema, table string, v interface{}, hashValues AttributeList, getAttributes AttributeList) error {
	return c.opRequest(operation{
		Operation:     OP_SEARCH_BY_HASH,
		Schema:        schema,
		Table:         table,
		HashValues:    hashValues,
		GetAttributes: getAttributes,
	}, &v)
}

// SearchByValue fetches records based on the value of an attribute
// Wilcards are allowed in `searchValue`
func (c *Client) SearchByValue(schema, table string, v interface{}, searchAttribute Attribute, searchValue interface{}, getAttributes AttributeList) error {
	return c.opRequest(operation{
		Operation:       OP_SEARCH_BY_VALUE,
		Schema:          schema,
		Table:           table,
		SearchAttribute: searchAttribute,
		SearchValue:     searchValue,
		GetAttributes:   getAttributes,
	}, &v)
}
