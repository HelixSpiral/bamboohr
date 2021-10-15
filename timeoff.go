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

// ListTimeOffRequests lists the timeoff requests
func (b *Client) ListTimeOffRequests(start, end time.Time, args map[string]interface{}) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/time_off/requests/?start=%s&end=%s", b.APIEndpoint, start.Format("2006-01-02"), end.Format("2006-01-02"))

	if id, ok := args["id"]; ok {
		endpointURL += fmt.Sprintf("&id=%s", id)
	}

	if action, ok := args["action"]; ok {
		endpointURL += fmt.Sprintf("&action=%s", action)
	}

	if empId, ok := args["employeeid"]; ok {
		endpointURL += fmt.Sprintf("&employeeId=%s", empId)
	}

	if tType, ok := args["type"]; ok {
		endpointURL += fmt.Sprintf("&type=%s", tType)
	}

	if status, ok := args["status"]; ok {
		endpointURL += fmt.Sprintf("&status=%s", status)
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
