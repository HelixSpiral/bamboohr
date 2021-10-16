package bamboohr

import (
	"fmt"
	"time"
)

// ListTableRowsForEmployee lists table rows for a given employee
func (b *Client) ListTableRowsForEmployee(id, table string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/%s/tables/%s", b.APIEndpoint, id, table)

	return b.getRequest(endpointURL)
}

// ListUpdatedEmployeeTabledata lists data updated since the provided time
func (b *Client) ListUpdatedEmployeeTabledata(table string, since time.Time) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employees/changed/tables/%s?since=%s", b.APIEndpoint, table, since.Format("2006-01-02"))

	return b.getRequest(endpointURL)
}
