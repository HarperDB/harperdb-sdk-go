package harperdb

import (
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	DEFAULT_ENDPOINT = "http://localhost:9925"
	DEFAULT_USERNAME = "HDB_ADMIN"
	DEFAULT_PASSWORD = "password"
)

var (
	c *Client
)

func init() {
	c = createClient()
}

func createClient() *Client {
	return NewClient(DEFAULT_ENDPOINT, DEFAULT_USERNAME, DEFAULT_PASSWORD)
}

func wait() {
	// TODO Clarify with Harper team why this is necessary
	time.Sleep(50 * time.Millisecond)
}

func randomID() string {
	return fmt.Sprintf("id-%s", uuid.NewV4().String())
}

func TestNewClient(t *testing.T) {
	_, err := c.RegistrationInfo()
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetFingerprint(t *testing.T) {
	_, err := c.GetFingerprint()
	if err != nil {
		t.Fatal(err)
	}

}
