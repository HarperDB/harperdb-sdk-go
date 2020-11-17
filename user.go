package harperdb

type User struct {
	Record
	Active   bool   `json:"active"`
	Role     Role   `json:"role"`
	Username string `json:"username"`
}

func (c *Client) AddUser(username, password, roleID string, active bool) error {
	return c.opRequest(operation{
		Operation: OP_ADD_USER,
		Role:      roleID,
		Username:  username,
		Password:  password,
		Active:    &active,
	}, nil)
}

func (c *Client) AlterUser(username, password, roleID string, active bool) error {
	return c.opRequest(operation{
		Operation: OP_ALTER_USER,
		Role:      roleID,
		Username:  username,
		Password:  password,
		Active:    &active,
	}, nil)
}

// DropUser deletes a user.
// Note: this operation is idempotent, it will not throw an error
// if the user doesn't exist
func (c *Client) DropUser(username string) error {
	return c.opRequest(operation{
		Operation: OP_DROP_USER,
		Username:  username,
	}, nil)
}

// UserInfo returns the current user executing this operation
func (c *Client) UserInfo() (User, error) {
	var user User
	err := c.opRequest(operation{
		Operation: OP_USER_INFO,
	}, &user)
	return user, err
}

func (c *Client) ListUsers() ([]User, error) {
	var users []User
	err := c.opRequest(operation{
		Operation: OP_LIST_USERS,
	}, &users)
	return users, err
}
