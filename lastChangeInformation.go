package bamboohr

import (
	"fmt"
	"net/http"
	"time"
)

// ListAllUpdatedEmployeeIds lists all the updated employee IDs
func (b *Client) ListAllUpdatedEmployeeIds(since time.Time, tType string) ([]byte, error) {
	switch tType {
	case "inserted":
	case "updated":
	case "deleted":
	default:
		return []byte{}, fmt.Errorf("type must be one of the following: inserted, updated, deleted")
	}
	endpointURL := fmt.Sprintf("%s/employees/changed?since=%s&type=%s", b.APIEndpoint, since.Format(time.RFC3339), tType)

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := b.sendRequest(req)
	if err != nil {
		return []byte{}, err
	}

	return resp, nil
}
