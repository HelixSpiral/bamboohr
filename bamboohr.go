package bamboohr

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client holds the client information for our BambooHR API calls
type Client struct {
	APIEndpoint string
	APIKey      string
	HTTPClient  *http.Client
}

// BBHRNewInput is used as input to the New() function
type BBHRNewInput struct {
	Version string
	Company string
	APIKey  string
}

// New creates a new BambooHR API Client and returns it to the caller
func New(args BBHRNewInput) (*Client, error) {
	newClient := &Client{}

	// Default version to v1
	if args.Version == "" {
		args.Version = "v1"
	}

	if args.Company == "" {
		return newClient, fmt.Errorf("missing arg: Company")
	}

	if args.APIKey == "" {
		return newClient, fmt.Errorf("missing arg: APIKey")
	} else {
		newClient.APIKey = args.APIKey
	}

	newClient.APIEndpoint = fmt.Sprintf("https://api.bamboohr.com/api/gateway.php/%s/%s", args.Company, args.Version)
	newClient.HTTPClient = http.DefaultClient

	return newClient, nil
}

// sendRequest sends the actual http request to BBHR
func (b *Client) sendRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("ACCEPT", "application/json")
	req.SetBasicAuth(b.APIKey, "x")

	resp, err := b.HTTPClient.Do(req)
	if resp != nil {
		defer resp.Body.Close() // This panics if body is nil, so we have to have the check first.
	}
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("status code not 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}

// getRequest abstraction
func (b *Client) getRequest(endpointURL string) ([]byte, error) {
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
