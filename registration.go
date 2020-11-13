package harperdb

import (
	"log"
	"time"
)

type RegistrationInfoResponse struct {
	Registered            bool      `json:"registered"`
	Version               string    `json:"version"`
	StorageType           string    `json:"storage_type"`
	RAMAllocation         int       `json:"ram_allocation"`
	LicenseExpirationDate time.Time `json:"license_expiration_date"`
}

func (c *Client) RegistrationInfo() (*RegistrationInfoResponse, error) {
	result := RegistrationInfoResponse{}

	resp, err := c.opRequest(operation{
		Operation: "registration_info",
	}, &result)

	if err != nil {
		return nil, err
	}

	log.Printf("RegistrationInfo: %d", resp.StatusCode())

	return &result, nil
}

type GetFingerprintResponse struct {
	Message string `json:"message"`
}

func (c *Client) GetFingerprint() (*GetFingerprintResponse, error) {
	result := GetFingerprintResponse{}

	resp, err := c.opRequest(operation{
		Operation: "get_fingerprint",
	}, &result)

	if err != nil {
		return nil, err
	}

	log.Printf("GetFingerprintInfo: %d", resp.StatusCode())

	return &result, nil
}
