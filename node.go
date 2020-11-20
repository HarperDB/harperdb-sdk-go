package harperdb

type Subscription struct {
	Channel   string `json:"channel"`
	Subscribe bool   `json:"subscribe"`
	Publish   bool   `json:"publish"`
}

type Connection struct {
	ID            string         `json:"id"`
	HostAddress   string         `json:"host_address"`
	HostPort      int            `json:"host_port"`
	State         string         `json:"state"`
	NodeName      string         `json:"node_name"`
	Subscriptions []Subscription `json:"subscriptions"`
}

type ClusterStatusResponse struct {
	IsEnabled bool        `json:"is_enabled"`
	NodeName  interface{} `json:"node_name"` // (sic) this is an int if cluster is not enabled
	Status    struct {
		ID                  string       `json:"id"`
		Type                string       `json:"type"`
		OutboundConnections []Connection `json:"outbound_connections"`
		InboundConnections  []Connection `json:"inbound_connections"`
	} `json:"status"`
}

func (c *Client) AddNode(name, host string, port int, subscriptions []Subscription) error {
	return c.opRequest(operation{
		Operation:     OP_ADD_NODE,
		Name:          name,
		Host:          host,
		Port:          port,
		Subscriptions: subscriptions,
	}, nil)
}

func (c *Client) UpdateNode(name, host string, port int, subscriptions []Subscription) error {
	return c.opRequest(operation{
		Operation:     OP_UPDATE_NODE,
		Name:          name,
		Host:          host,
		Port:          port,
		Subscriptions: subscriptions,
	}, nil)
}

func (c *Client) RemoveNode(name string) error {
	return c.opRequest(operation{
		Operation: OP_REMOVE_NODE,
		Name:      name,
	}, nil)
}

func (c *Client) ClusterStatus() (*ClusterStatusResponse, error) {
	var result ClusterStatusResponse
	err := c.opRequest(operation{
		Operation: OP_CLUSTER_STATUS,
	}, &result)
	return &result, err
}
