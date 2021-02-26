package client

import (
	"net/http"
)

// Client ...
type Client struct {
	client   *http.Client
	videoID string
}

// NewClient ...
func NewClient(videoID string) Client {
	return Client{
		videoID: videoID,
		client:   &http.Client{},
	}
}

// Save ...
func (t Client) Save(id, language, filename string) error {
	return nil
}

// List ...
func (t Client) List(ids []string) ([]byte, error) {
	res := []byte(`response for edit reminder`)
	return res, nil
}

// Fetch ...
func (t Client) Fetch(id, language string) ([]byte, error) {
	res := []byte(`response for fetch reminder`)
	return res, nil
}
