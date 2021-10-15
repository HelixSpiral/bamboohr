package bamboohr

import (
	"fmt"
	"net/http"
	"time"
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

// ListTimeOffPoliciesForEmployee lists the time off policies for a specific employee
func (b *Client) ListTimeOffPoliciesForEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/time_off/policies", b.APIEndpoint, id)

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

// EstimateFutureTimeOffBalance estimates the future time off balance for an employee
func (b *Client) EstimateFutureTimeOffBalance(id string, end time.Time) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/time_off/calculator?end=%s", b.APIEndpoint, id, end.String())

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

// ListWhosOut lists the employees out for the specified dates
func (b *Client) ListWhosOut(args map[string]interface{}) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/time_off/whos_out/", b.APIEndpoint)
	if startDate, ok := args["start"]; !ok {
		endpointURL += fmt.Sprintf("?start=%s", startDate)
	}

	if endDate, ok := args["end"]; !ok {
		endpointURL += fmt.Sprintf("&end=%s", endDate)
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
