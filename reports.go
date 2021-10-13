package bamboohr

import (
	"fmt"
	"net/http"
)

// GetCompanyReport gets a company report
func (b *Client) GetCompanyReport(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/reports/%s", b.APIEndpoint, id)

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
