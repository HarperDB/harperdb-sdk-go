package harperdb

import (
	"fmt"
	"testing"
)

func TestListUsers(t *testing.T) {
	users, err := c.ListUsers()
	if err != nil {
		t.Fatal(err)
	}

	// We expect at least to have the HDB_ADMIN
	var found bool
	for _, user := range users {
		if user.Username == "HDB_ADMIN" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("expected to find HDB_ADMIN user")
	}
}

func TestAddUser(t *testing.T) {
	superUser, err := findRole(SUPER_USER)
	if err != nil {
		t.Fatal(err)
	}

	testUser := randomID()
	err = c.AddUser(testUser, randomID(), superUser.ID, true)
	if err != nil {
		t.Fatal(err)
	}

	user, err := findUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal(fmt.Sprintf("expected to find user %s", testUser))
	}

	err = c.DropUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlterUser(t *testing.T) {
	superUser, err := findRole(SUPER_USER)
	if err != nil {
		t.Fatal(err)
	}

	testUser := randomID()
	err = c.AddUser(testUser, randomID(), superUser.ID, true)
	if err != nil {
		t.Fatal(err)
	}

	// set user to inactive
	err = c.AlterUser(testUser, randomID(), superUser.ID, false)
	if err != nil {
		t.Fatal(err)
	}

	user, err := findUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal(fmt.Sprintf("expected to find user %s", testUser))
	}
	if user.Active {
		t.Fatal("did not expect user to be active")
	}

	err = c.DropUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserInfo(t *testing.T) {
	user, err := c.UserInfo()
	if err != nil {
		t.Fatal(err)
	}

	if user.Username != "HDB_ADMIN" {
		t.Fatal(fmt.Errorf("expected user to be HDB_ADMIN"))
	}
}

func findUser(username string) (*User, error) {
	users, err := c.ListUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, nil
}
