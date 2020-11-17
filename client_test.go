package harperdb

import (
	"fmt"
	"strings"
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
	// According to HarperDB some operations are asyncronously penetrated between
	// processes, so while the objects are created they are not immediately visible
	time.Sleep(200 * time.Millisecond)
}

func randomID() string {
	return fmt.Sprintf("id_%s", strings.ReplaceAll(uuid.NewV4().String(), "-", "_"))
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
