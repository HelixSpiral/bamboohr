package bamboohr

import (
	"fmt"
)

// GetCompanyReport gets a company report
func (b *Client) GetCompanyReport(args map[string]string) ([]byte, error) {
	if _, ok := args["id"]; !ok {
		return []byte{}, fmt.Errorf("missing arg: id")
	}

	endpointURL := fmt.Sprintf("%s/reports/%s", b.APIEndpoint, args["id"])

	if _, ok := args["format"]; !ok {
		return []byte{}, fmt.Errorf("missing arg: format")
	}

	switch args["format"] {
	case "csv":
	case "pdf":
	case "xls":
	case "xml":
	case "json":
	default:
		return []byte{}, fmt.Errorf("arg 'format' must be one of: csv, pdf, xls, xml, json")
	}

	endpointURL += fmt.Sprintf("&format=%s", args["format"])

	if _, ok := args["fd"]; ok {
		endpointURL += fmt.Sprintf("&fd=%s", args["fd"])
	}

	return b.getRequest(endpointURL)
}
