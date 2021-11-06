package bamboohr

import (
	"fmt"
	"time"
)

// ListTimeOffTypes lists the types of time off in your account
func (b *Client) ListTimeOffTypes(request bool) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/time_off/types", b.APIEndpoint)

	if request {
		endpointURL += "/?mode=request"
	}

	return b.getRequest(endpointURL)
}

// ListTimeOffPolicies lists the time off policies
func (b *Client) ListTimeOffPolicies() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/meta/time_off/policies", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// BBHRListTimeOffRequestsInput is used as input for the ListTimeOffRequests function
type BBHRListTimeOffRequestsInput struct {
	Start time.Time
	End   time.Time

	TimeOffID  string
	EmployeeID string

	Action string
	Status string
	Type   string
}

// ListTimeOffRequests lists the timeoff requests
func (b *Client) ListTimeOffRequests(args BBHRListTimeOffRequestsInput) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/time_off/requests/?start=%s&end=%s", b.APIEndpoint, args.Start.Format("2006-01-02"), args.End.Format("2006-01-02"))

	if args.TimeOffID != "" {
		endpointURL += fmt.Sprintf("&id=%s", args.TimeOffID)
	}

	if args.Action != "" {
		endpointURL += fmt.Sprintf("&action=%s", args.Action)
	}

	if args.EmployeeID != "" {
		endpointURL += fmt.Sprintf("&employeeId=%s", args.EmployeeID)
	}

	if args.Type != "" {
		endpointURL += fmt.Sprintf("&type=%s", args.Type)
	}

	if args.Status != "" {
		endpointURL += fmt.Sprintf("&status=%s", args.Status)
	}

	return b.getRequest(endpointURL)
}

// ListTimeOffPoliciesForEmployee lists the time off policies for a specific employee
func (b *Client) ListTimeOffPoliciesForEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/time_off/policies", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// EstimateFutureTimeOffBalance estimates the future time off balance for an employee
func (b *Client) EstimateFutureTimeOffBalance(id string, end time.Time) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/time_off/calculator?end=%s", b.APIEndpoint, id, end.String())

	return b.getRequest(endpointURL)
}

// ListWhosOut lists the employees out for the specified dates
func (b *Client) ListWhosOut(args map[string]time.Time) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/time_off/whos_out/", b.APIEndpoint)
	if startDate, ok := args["start"]; ok {
		endpointURL += fmt.Sprintf("?start=%s", startDate.Format("2006-01-02"))
	}

	if endDate, ok := args["end"]; ok {
		endpointURL += fmt.Sprintf("&end=%s", endDate.Format("2006-01-02"))
	}

	return b.getRequest(endpointURL)
}
