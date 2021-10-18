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

// New creates a new BambooHR API Client and returns it to the caller
func New(args map[string]interface{}) (*Client, error) {
	newClient := &Client{}

	if _, ok := args["version"]; !ok {
		args["version"] = "v1"
	}

	if _, ok := args["company"]; !ok {
		return newClient, fmt.Errorf("missing arg: company")
	}

	if _, ok := args["apikey"]; !ok {
		return newClient, fmt.Errorf("missing arg: apikey")
	} else {
		newClient.APIKey = args["apikey"].(string)
	}

	newClient.APIEndpoint = fmt.Sprintf("https://api.bamboohr.com/api/gateway.php/%s/%s", args["company"], args["version"])

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
