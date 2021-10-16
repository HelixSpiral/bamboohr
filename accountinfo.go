package bamboohr

import (
	"fmt"
)

// ListFields lists all the metadata fields your account has
func (b *Client) ListFields() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/fields", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListTabularFields lists all the tables for your account
func (b *Client) ListTabularFields() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/tables", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListListFields lists all the list fields your account has
func (b *Client) ListListFields() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/lists", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListUsers lists all the users in your account
func (b *Client) ListUsers() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/users", b.APIEndpoint)

	return b.getRequest(endpointURL)
}
