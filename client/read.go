package client

import (
	"bytes"

	"github.com/AirHelp/treasury/types"
	"github.com/AirHelp/treasury/utils"
)

// Read returns decrypted secret for given key
func (c *Client) Read(key string) (*Secret, error) {
	if err := utils.ValidateInputKey(key); err != nil {
		return nil, err
	}

	s3objectInput := &types.GetObjectInput{
		Key: key,
	}
	s3object, err := c.Backend.GetObject(s3objectInput)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(s3object.Body)

	secret := &Secret{
		Key:   key,
		Value: buf.String(),
	}
	return secret, nil
}

// ReadValue returns secret as a string for given key.
func (c *Client) ReadValue(key string) (string, error) {
	secret, err := c.Read(key)
	if err != nil {
		return "", err
	}
	return secret.Value, nil
}

// ReadGroup returns list of secrets for given key prefix
func (c *Client) ReadGroup(keyPrefix string) ([]*Secret, error) {
	if err := utils.ValidateInputKeyPattern(keyPrefix); err != nil {
		return nil, err
	}
	params := &types.GetObjectsInput{
		Prefix: keyPrefix,
	}
	resp, err := c.Backend.GetObjects(params)

	if err != nil {
		return nil, err
	}

	var secrets []*Secret

	for key, value := range resp.Secrets {
		secret := &Secret{
			Key:   key,
			Value: value,
		}
		secrets = append(secrets, secret)
	}
	return secrets, nil
}
