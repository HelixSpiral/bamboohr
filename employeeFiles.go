package bamboohr

import (
	"fmt"
)

// ListFilesByEmployee will list all the files a given employee has in their account
func (b *Client) ListFilesByEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/files/view", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetFileFromEmployee will get a specified file from an employee
func (b *Client) GetFileFromEmployee(empID, fileID string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/files/%s", b.APIEndpoint, empID, fileID)

	return b.getRequest(endpointURL)
}
