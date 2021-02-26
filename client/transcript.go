package client

import (
	"errors"
	"net/http"

	"github.com/SirNoob97/yt-transcripts/transcript"
)

// Client ...
type Client struct {
	client  *http.Client
	videoID string
}

// NewClient ...
func NewClient(videoID string) Client {
	return Client{
		videoID: videoID,
		client:  &http.Client{},
	}
}

// Save ...
func (t Client) Save(id, language, filename string) error {
	return nil
}

// List ...
func (t Client) List(ids []string) ([]string, error) {
	return transcript.ListTranscripts(ids, t.client)
}

// Fetch ...
func (t Client) Fetch(id, language string) ([]string, error) {
	tr := transcript.FetchTranscript(id, language, t.client)
	if len(tr.Text) < 0 {
		return "", errors.New("Captions Not Avalible")
	}
	return tr.Text, nil
}
