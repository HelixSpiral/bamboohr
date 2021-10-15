package bamboohr

import (
	"fmt"
)

// GetCompanyReport gets a company report
func (b *Client) GetCompanyReport(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/reports/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}
