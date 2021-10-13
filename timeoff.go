package bamboohr

import (
	"fmt"
	"net/http"
)

// ListTimeOffTypes lists the types of time off in your account
func (b *Client) ListTimeOffTypes(request bool) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/time_off/types", b.APIEndpoint)

	if request {
		endpointURL += "/?mode=request"
	}

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

// ListTimeOffPolicies lists the time off policies
func (b *Client) ListTimeOffPolicies() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/time_off/policies", b.APIEndpoint)

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
