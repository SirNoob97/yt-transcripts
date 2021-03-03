package cli

import (
	"errors"
	"fmt"
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
	languages, err := t.fetcher.List(id)
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0, len(languages))
	ret = append(ret, fmt.Sprintf("Available Transcripts of %s", id))
	i := 1
	for name, lan := range languages {
		ret = append(ret, fmt.Sprintf("Transcript #%d %s(%s)", i, name, lan))
		i++
	}
	return ret, nil
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
