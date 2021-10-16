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

// ListAvailableGoalSharingOptions lists the employees with whom the specified employees goals can be shared
func (b *Client) ListAvailableGoalSharingOptions(id string, args map[string]string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/performance/employees/%s/goals/sharedOptions", b.APIEndpoint, id)

	if searchOptions, ok := args["search"]; ok {
		endpointURL += fmt.Sprintf("&search=%s", searchOptions)
	}

	if limit, ok := args["limit"]; ok {
		endpointURL += fmt.Sprintf("&limit=%s", limit)
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
