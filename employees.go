package bamboohr

import (
	"fmt"
)

// ListEmployees lists all the employees in your company
func (b *Client) ListEmployees() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/%s", b.APIEndpoint, "employees/directory")

	return b.getRequest(endpointURL)
}

// BBHRGetEmployeeInput is used as an input struct for the GetEmployee function
type BBHRGetEmployeeInput struct {
	EmployeeID string
	Fields     string
}

// GetEmployee gets a specific employee
func (b *Client) GetEmployee(args BBHRGetEmployeeInput) ([]byte, error) {
	if args.EmployeeID == "" {
		return []byte{}, fmt.Errorf("missing field: EmployeeID")
	}
	if args.Fields == "" {
		args.Fields = "firstName,lastName"
	}

	endpointURL := fmt.Sprintf("%s/employees/%s/?fields=%s", b.APIEndpoint, args.EmployeeID, args.Fields)

	return b.getRequest(endpointURL)
}
