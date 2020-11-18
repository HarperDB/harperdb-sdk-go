package harperdb

import "testing"

func TestAddAndRemoveNode(t *testing.T) {
	// Try to remove node first in case previous tests failed
	err := c.RemoveNode("node2")
	if err != nil {
		t.Fatal(err)
	}

	err = c.AddNode("node2", "127.0.0.1", 1112, []Subscription{
		{
			Channel:   "dev:dog",
			Publish:   true,
			Subscribe: false,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = c.UpdateNode("node2", "127.0.0.1", 1112, []Subscription{
		{
			Channel:   "dev:dog",
			Publish:   true,
			Subscribe: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = c.RemoveNode("node2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCluserStatus(t *testing.T) {
	_, err := c.ClusterStatus()
	if err != nil {
		t.Fatal(err)
	}

	// TODO Here we could do a little more testing
}
