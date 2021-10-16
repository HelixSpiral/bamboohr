package bamboohr

import "fmt"

// GetTimeTrackingRecord gets a specific time record
func (b *Client) GetTimeTrackingRecord(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/timetracking/record/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}
