package bamboohr

import (
	"fmt"
	"net/http"
)

// ListTrainingTypes lists the types of training your company has
func (b *Client) ListTrainingTypes() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/type", b.APIEndpoint)

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

// ListTrainingCategories lists the training categories your company has
func (b *Client) ListTrainingCategories() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/category", b.APIEndpoint)

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

// ListEmployeeTrainingsByType lists the training of an employee for a specific type
func (b *Client) ListEmployeeTrainingsByType(id, tType string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/record/employee/%s/type/%s", b.APIEndpoint, id, tType)

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

// ListEmployeeTrainings lists the trainings by an employee
func (b *Client) ListEmployeeTrainings(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/record/employee/%s", b.APIEndpoint, id)

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
