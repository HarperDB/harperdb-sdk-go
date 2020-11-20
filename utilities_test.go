package harperdb

import "testing"

func TestSystemStatus(t *testing.T) {
	guest := "guest"

	// delete guest user and role
	c.DropUser(guest)

	// try with standard super user
	_, err := c.SystemInformation()
	if err != nil {
		t.Fatal(err)
	}

	// try with non-super user
	p := Permission{}
	p.SetSuperUser(false)
	p.SetClusterUser(false)
	role, err := c.AddRole(guest, p)
	if err != nil {
		t.Fatal(err)
	}

	err = c.AddUser(guest, guest, role.ID, true)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = c.DropUser(guest)
		_ = c.DropRole(role.ID)
	}()

	guestClient := NewClient("http://localhost:9925", guest, guest)
	_, err = guestClient.SystemInformation()
	if e, ok := err.(*OperationError); ok && e.IsNotAuthorizedError() {
		return
	}

	t.Fatal("expected SystemInfo call with guest user to fail")
}
