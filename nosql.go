package harperdb

/*


{
  "message": "inserted 2 of 2 records",
  "skipped_hashes": [],
  "inserted_hashes": [
    "2a35ecf8-26e1-4526-8e49-d2a78ba53433",
    "f7b9dd17-e79f-4e58-afe6-fcaa5d39237b"
  ]
}

*/

type AffectedResponse struct {
	MessageResponse
	SkippedHashes  []string `json:"skipped_hashes"`
	InsertedHashes []string `json:"inserted_hashes"`
	UpdatedHashes  []string `json:"update_hashes"` // (sic) not updated_hashes
	DeletedHashes  []string `json:"deleted_hashes"`
}

// Insert inserts one or more JSON objects into a table
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Insert(schema, table string, records interface{}) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation: OP_INSERT,
		Schema:    schema,
		Table:     table,
		Records:   records,
	}, &result)
	return &result, err
}

// Update updates one or more JSON objects in a table.
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Update(schema, table string, records interface{}) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation: OP_UPDATE,
		Schema:    schema,
		Table:     table,
		Records:   records,
	}, &result)
	return &result, err
}

// Delete delete one or more JSON objects from a table.
// hashValues must be an array of slice
func (c *Client) Delete(schema, table string, hashValues AttributeList) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation:  OP_DELETE,
		Schema:     schema,
		Table:      table,
		HashValues: hashValues,
	}, &result)
	return &result, err
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
