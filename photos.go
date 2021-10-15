package bamboohr

import (
	"fmt"
	"net/http"
)

// GetEmployeePhoto
func (b *Client) GetEmployeePhoto(id, size string) ([]byte, error) {
	switch size {
	case "original":
	case "large":
	case "medium":
	case "xs":
	case "tiny":
	default:
		return []byte{}, fmt.Errorf("size must be one of: original, large, medium, xs, tiny")
	}
	endpointURL := fmt.Sprintf("%s/employees/%s/photo/%s", b.APIEndpoint, id, size)

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
