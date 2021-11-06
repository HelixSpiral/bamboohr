package bamboohr

import "fmt"

// GetGoalStatusCount gets the number of goals per status for an employee
func (b *Client) GetGoalStatusCount(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals/filters", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetCanCreateGoal returns if the API user can create a goal for the employee
func (b *Client) GetCanCreateGoal(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals/canCreateGoals", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListGoalsForEmployee lists the goals for an employee
func (b *Client) ListGoalsForEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// BBHRListAvailableGoalSharingOptionsInput is used as input for the ListAvailableGoalSharingOptions function
type BBHRListAvailableGoalSharingOptionsInput struct {
	EmployeeID string
	Search     string
	Limit      string
}

// ListAvailableGoalSharingOptions lists the employees with whom the specified employees goals can be shared
func (b *Client) ListAvailableGoalSharingOptions(args BBHRListAvailableGoalSharingOptionsInput) ([]byte, error) {
	if args.EmployeeID == "" {
		return []byte{}, fmt.Errorf("missing arg: EmployeeID")
	}

	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals/sharedOptions", b.APIEndpoint, args.EmployeeID)

	if args.Search != "" {
		endpointURL += fmt.Sprintf("&search=%s", args.Search)
	}

	if args.Limit != "" {
		endpointURL += fmt.Sprintf("&limit=%s", args.Limit)
	}

	return b.getRequest(endpointURL)
}

// GetAlignableGoalOptions gets the alignable goal options for an employee
func (b *Client) GetAlignableGoalOptions(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals/alignmentOptions", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetGoalComments gets the comments for a specific goal
func (b *Client) GetGoalComments(id, goalid string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals/%s/comments", b.APIEndpoint, id, goalid)

	return b.getRequest(endpointURL)
}
