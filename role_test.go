package harperdb

import (
	"fmt"
	"testing"
)

const (
	CLUSTER_USER = "cluster_user"
	SUPER_USER   = "super_user"
)

func TestListRoles(t *testing.T) {
	roles, err := c.ListRoles()
	if err != nil {
		t.Fatal(err)
	}

	// We should find two roles, super_user and cluster_user
	var foundSU, foundCU bool
	for _, role := range roles {
		switch role.Role {
		case CLUSTER_USER:
			foundCU = true
		case SUPER_USER:
			foundSU = true
		}
	}

	if !(foundSU && foundCU) {
		t.Fatal("did not find super_user or cluster_user role")
	}
}

func TestDropAndAddRole(t *testing.T) {
	foundCU, err := findRole(CLUSTER_USER)
	if err != nil {
		t.Fatal(err)
	}

	if foundCU == nil {
		t.Fatal("did not find cluster user role")
	}

	t.Log(fmt.Sprintf("Found cluster user role id: %s", foundCU.ID))
	if err := c.DropRole(foundCU.ID); err != nil {
		t.Fatal(err)
	}

	role := Permission{}
	role.SetClusterUser(true)
	newRole, err := c.AddRole(CLUSTER_USER, role)
	if err != nil {
		t.Fatal(err)
	}
	if newRole.Role != CLUSTER_USER {
		t.Fatal(fmt.Errorf("expected new role named %s", CLUSTER_USER))
	}
}

func TestAlterRole(t *testing.T) {
	foundCU, err := findRole(CLUSTER_USER)
	if err != nil {
		t.Fatal(err)
	}

	if foundCU == nil {
		t.Fatal("did not find cluster user role")
	}

	newName := foundCU.Role + "new"
	newRole, err := c.AlterRole(foundCU.ID, newName, foundCU.Permission)
	if err != nil {
		t.Fatal(err)
	}
	if newRole.Role != newName {
		t.Fatal(fmt.Errorf("expected altered role to have name %s", newName))
	}

	// Change name back to original
	_, err = c.AlterRole(foundCU.ID, foundCU.Role, foundCU.Permission)
	if err != nil {
		t.Fatal(err)
	}
}

func findRole(name string) (*Role, error) {
	roles, err := c.ListRoles()
	if err != nil {
		return nil, err
	}

	var found *Role
	for _, role := range roles {
		if role.Role == name {
			found = &role
			break
		}
	}
	return found, nil
}
