package cli

import (
	"errors"
	"html"
	"log"
	"os"
	"strings"

	"github.com/SirNoob97/yt-transcripts/transcript"
)

// FetcherClient ...
type FetcherClient struct {
	fetcher transcript.Fetcher
}

// NewFetcherClient ...
func NewFetcherClient(fetcher transcript.Fetcher) FetcherClient {
	return FetcherClient{fetcher: fetcher}
}

// Save ...
func (t FetcherClient) Save(id, language, filename string) error {
	tr, err := t.Fetch(id, language)
	if err != nil {
		return err
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
func (t FetcherClient) List(id string) ([]string, error) {
	return t.fetcher.List(id)
}

// Fetch ...
func (t FetcherClient) Fetch(id, language string) (string, error) {
	if language == "" {
		language = getSystemLanguage()
	}

	tr := t.fetcher.Fetch(id, language)
	if len(tr.Text) == 0 {
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
