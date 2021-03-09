package client

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Requester ...
type Requester interface {
	DoGetRequest(url string) ([]byte, error)
}

// HTTPClient ...
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient ...
func NewHTTPClient() HTTPClient {
	return HTTPClient{client: &http.Client{}}
}

// DoGetRequest ...
func (c HTTPClient) DoGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
