package bamboohr

import (
	"fmt"
)

// ListTrainingTypes lists the types of training your company has
func (b *Client) ListTrainingTypes() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/type", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListTrainingCategories lists the training categories your company has
func (b *Client) ListTrainingCategories() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/category", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListEmployeeTrainingsByType lists the training of an employee for a specific type
func (b *Client) ListEmployeeTrainingsByType(id, tType string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/record/employee/%s/type/%s", b.APIEndpoint, id, tType)

	return b.getRequest(endpointURL)
}

// ListEmployeeTrainings lists the trainings by an employee
func (b *Client) ListEmployeeTrainings(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/training/record/employee/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}
