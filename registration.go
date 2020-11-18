package harperdb

import (
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

	err := c.opRequest(operation{
		Operation: OP_REGISTRATION_INFO,
	}, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

type GetFingerprintResponse struct {
	Message string `json:"message"`
}

func (c *Client) GetFingerprint() (*GetFingerprintResponse, error) {
	result := GetFingerprintResponse{}

	err := c.opRequest(operation{
		Operation: OP_GET_FINGERPRINT,
	}, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetLicense(key, company string) error {
	return c.opRequest(operation{
		Operation: OP_SET_LICENSE,
		Company:   company,
		Key:       key,
	}, nil)
}
