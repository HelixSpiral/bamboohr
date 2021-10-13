package bamboohr

import (
	"fmt"
	"net/http"
)

// ListCompanyFiles lists all the files in your company
func (b *Client) ListCompanyFiles() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/files/view", b.APIEndpoint)

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

// GetCompanyFile gets a specific company file
func (b *Client) GetCompanyFile(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/files/%s", b.APIEndpoint, id)

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
