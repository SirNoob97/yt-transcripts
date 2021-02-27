package client

import (
	"errors"
	"html"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SirNoob97/yt-transcripts/transcript"
)

// Client ...
type Client struct {
	client  *http.Client
	videoID string
}

// NewClient ...
func NewClient() Client {
	return Client{
		client: &http.Client{},
	}
}

// Save ...
func (t Client) Save(id, language, filename string) error {
	tr, err := t.Fetch(id, language)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(tr)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// List ...
func (t Client) List(id string) ([]string, error) {
	return transcript.ListTranscripts(id, t.client)
}

// Fetch ...
func (t Client) Fetch(id, language string) (string, error) {
	if language == "" {
		language = getSystemLanguage()
	}

	tr := transcript.FetchTranscript(id, language, t.client)
	if len(tr.Text) < 0 {
		return "", errors.New("Captions Not Avalible")
	}

	return html.UnescapeString(strings.Join(tr.Text, "\n")), nil
}

func getSystemLanguage() string {
	str := os.Getenv("LANGUAGE")
	if str != "" {
		return strings.Split(str, ":")[1]
	}
	return "en"
}
