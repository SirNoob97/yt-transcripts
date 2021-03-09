package cli

import (
	"errors"
	"os"
	"strings"
	"testing"

	mocks "github.com/SirNoob97/yt-transcripts/mocks/transcript"
	"github.com/SirNoob97/yt-transcripts/transcript"
)

func TestNewClient(t *testing.T) {
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	if (client == FetcherClient{}) {
		t.Fatal("Expected an non-empty client")
	}
}

func TestSaveFetch(t *testing.T) {
	const id, lan, out = "ID", "LANGUAGE", "test/testSave"
	text := []string{"TEXT", "TEST"}
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	f.On("Fetch", id, lan).Return(transcript.Transcript{Text: text})

	err := client.Save(id, lan, out)

	if err != nil {
		t.Fatalf("Expected a nil err, got %v", err)
	}
}

func TestSaveFetchFailCase(t *testing.T) {
	const id, lan, out = "ID", "LANGUAGE", "test/testSave"
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	f.On("Fetch", id, lan).Return(transcript.Transcript{})

	err := client.Save(id, lan, out)

	if err == nil {
		t.Fatal("Expected an error message, got nil")
	}
}

func TestListFetch(t *testing.T) {
	const id, key, value = "ID", "KEY", "VALUE"
	res := map[string]string{key:value}
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	f.On("List", id).Return(res, nil)

	arr, err := client.List(id)
	if err != nil {
		t.Fatalf("Expected a nil error, got %v", err)
	}

	if len(arr) == 0 {
		t.Fatal("Expected a non-empty strings array")
	}
}

func TestListFetchFailCase(t *testing.T) {
	const id = "ID"
	var res, errorMsg = map[string]string{}, errors.New("ERROR")
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	f.On("List", id).Return(res, errorMsg)

	arr, err := client.List(id)
	if err == nil {
		t.Fatalf("Expected an error message, got nil")
	}

	if len(arr) != 0 {
		t.Fatalf("Expected an empty strings array, got %v", arr)
	}
}

func TestFetchFetch(t *testing.T) {
	const id, lan = "ID", "LANGUAGE"
	text := []string{"TEXT", "TEST"}
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	f.On("Fetch", id, lan).Return(transcript.Transcript{Text: text})

	res, err := client.Fetch(id, lan)
	if err != nil {
		t.Fatalf("Expected a nil err, got %v", err)
	}

	count := 0
	for _, t := range text {
		if strings.Contains(res, t) {
			count++
		}
	}

	if count != len(text) {
		t.Fatalf("Expected an string with '%v', got an empty string", text)
	}
}

func TestFetchFetchFailCase(t *testing.T) {
	const id, lan = "ID", "LANGUAGE"
	f := new(mocks.Fetcher)
	client := NewFetcherClient(f)

	f.On("Fetch", id, lan).Return(transcript.Transcript{})

	res, err := client.Fetch(id, lan)
	if err == nil {
		t.Fatal("Expected an error message, got nil")
	}

	if len(res) != 0 {
		t.Fatalf("Expected an empty string, got %s", res)
	}
}

func TestGetSystemLanguage(t *testing.T) {
	const env = "LANGUAGE"
	value := []string{"TEST", "TEST"}
	err := os.Setenv(env, strings.Join(value, ":"))
	if err != nil {
		t.Fatalf("ENV error %v", err)
	}

	lan := getSystemLanguage()

	if lan != value[0] {
		t.Fatalf("Expected %s as system language, got %s", value[0], lan)
	}
}

func TestGetSystemLanguageDefaultCase(t *testing.T) {
	const env = "LANGUAGE"
	err := os.Setenv(env, "")
	if err != nil {
		t.Fatalf("ENV error %v", err)
	}

	lan := getSystemLanguage()

	if len(lan) == 0 {
		t.Fatalf("Expected a default system language, got nothing")
	}
}
