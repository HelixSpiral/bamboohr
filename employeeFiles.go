package bamboohr

import (
	"fmt"
)

// ListFilesByEmployee will list all the files a given employee has in their account
func (b *Client) ListFilesByEmployee(args map[string]interface{}) ([]byte, error) {
	if _, ok := args["employeeid"]; !ok {
		return []byte{}, fmt.Errorf("missing field: employeeid")
	}

	endpointURL := fmt.Sprintf("%s/%s/%s/files/view", b.APIEndpoint, "employees", args["employeeid"])

	return b.getRequest(endpointURL)
}

// GetFileFromEmployee will get a specified file from an employee
func (b *Client) GetFileFromEmployee(args map[string]interface{}) ([]byte, error) {
	if _, ok := args["employeeid"]; !ok {
		return []byte{}, fmt.Errorf("missing field: employeeid")
	}

	if _, ok := args["fileid"]; !ok {
		return []byte{}, fmt.Errorf("missing field: fileid")
	}

	endpointURL := fmt.Sprintf("%s/%s/%s/files/%s", b.APIEndpoint, "employees", args["employeeid"], args["fileid"])

	return b.getRequest(endpointURL)
}
