package bamboohr

import (
	"fmt"
)

// BBHRGetCompanayReportInput is used as input for the GetCompanyReport function
type BBHRGetCompanyReportInput struct {
	ReportID string
	Format   string
	FD       string
}

// GetCompanyReport gets a company report
func (b *Client) GetCompanyReport(args BBHRGetCompanyReportInput) ([]byte, error) {
	if args.ReportID == "" {
		return []byte{}, fmt.Errorf("missing arg: ReportID")
	}

	endpointURL := fmt.Sprintf("%s/reports/%s", b.APIEndpoint, args.ReportID)

	if args.Format == "" {
		return []byte{}, fmt.Errorf("missing arg: Format")
	}

	switch args.Format {
	case "csv":
	case "pdf":
	case "xls":
	case "xml":
	case "json":
	default:
		return []byte{}, fmt.Errorf("arg 'format' must be one of: csv, pdf, xls, xml, json")
	}

	endpointURL += fmt.Sprintf("&format=%s", args.Format)

	if args.FD != "" {
		endpointURL += fmt.Sprintf("&fd=%s", args.FD)
	}

	return b.getRequest(endpointURL)
}
