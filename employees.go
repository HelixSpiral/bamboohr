package bamboohr

import (
	"fmt"
)

// ListEmployees lists all the employees in your company
func (b *Client) ListEmployees() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/%s", b.APIEndpoint, "employees/directory")

	return b.getRequest(endpointURL)
}

// GetEmployee gets a specific employee
func (b *Client) GetEmployee(args map[string]interface{}) ([]byte, error) {
	if _, ok := args["employeeid"]; !ok {
		return []byte{}, fmt.Errorf("missing field: employeeid")
	}
	if _, ok := args["fields"]; !ok {
		args["fields"] = "firstName,lastName"
	}

	endpointURL := fmt.Sprintf("%s/%s/%s/?fields=%s", b.APIEndpoint, "employees", args["employeeid"], args["fields"])

	return b.getRequest(endpointURL)
}
