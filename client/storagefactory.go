package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateRecord - Create new A Record with CloudTribe client
func (c *Client) CreateRecord(sfModel StorageFactoryModel) (interface{}, error) {

	// Convert Request Data/Body to JSON
	rb, err := json.Marshal(sfModel)
	if err != nil {
		return nil, err
	}

	// Prepare the URL, Method and Payload fo the Client
	url := fmt.Sprintf("%s/storagefactory", c.SfcpEndpoint)
	method := "POST"
	payload := strings.NewReader(string(rb))

	// Create a new Request
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	// Add Proper Headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "Oracle")) //FIXME:
	req.Host = "sfcp-api.gcd.com"

	// Make the Request to create record
	res, err := c.CustomHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// Read the ResponseBody content
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Check Status Code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	// Preparing Response Body to return and convert it to Golang Map Object
	var responseBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &responseBody)
	if err != nil {
		return nil, err
	}

	// Get the ID of the resource from the Response Body
	ocid := responseBody["sfocid"]

	// Return ID of the record created
	return ocid, nil

}
