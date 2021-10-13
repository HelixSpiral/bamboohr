package bamboohr

import (
	"fmt"
	"net/http"
)

// ListFields lists all the metadata fields your account has
func (b *Client) ListFields() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/field", b.APIEndpoint)

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := b.sendRequest(req)
	if err != nil {
		return []byte{}, err
	}

	return resp, nil
}

// ListTabularFields lists all the tables for your account
func (b *Client) ListTabularFields() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/tables", b.APIEndpoint)

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := b.sendRequest(req)
	if err != nil {
		return []byte{}, err
	}

	return resp, nil
}

// ListListFields lists all the list fields your account has
func (b *Client) ListListFields() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/lists", b.APIEndpoint)

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := b.sendRequest(req)
	if err != nil {
		return []byte{}, err
	}

	return resp, nil
}

// ListUsers lists all the users in your account
func (b *Client) ListUsers() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/users", b.APIEndpoint)

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := b.sendRequest(req)
	if err != nil {
		return []byte{}, err
	}

	return resp, nil
}
