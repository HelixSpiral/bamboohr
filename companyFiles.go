package bamboohr

import (
	"fmt"
)

// ListCompanyFiles lists all the files in your company
func (b *Client) ListCompanyFiles() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/files/view", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetCompanyFile gets a specific company file
func (b *Client) GetCompanyFile(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/files/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}
