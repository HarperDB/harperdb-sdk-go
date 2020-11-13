package harperdb

import (
	"log"
	"testing"
)

const (
	DEFAULT_ENDPOINT = "http://localhost:9925"
	DEFAULT_USERNAME = "HDB_ADMIN"
	DEFAULT_PASSWORD = "password"
)

func createClient() *Client {
	return NewClient(DEFAULT_ENDPOINT, DEFAULT_USERNAME, DEFAULT_PASSWORD)
}
func TestNewClient(t *testing.T) {
	c := createClient()

	regInfo, err := c.RegistrationInfo()
	if err != nil {
		t.Fail()
	}

	log.Println(regInfo)
}

func TestGetFingerprint(t *testing.T) {
	c := createClient()

	fingerprint, err := c.GetFingerprint()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	log.Println(fingerprint)
}
