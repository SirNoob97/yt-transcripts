package client

import (
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	httpClient := NewHTTPClient()

	if (httpClient == HTTPClient{}) {
		t.Fatal("Expected a non empty HTTPClient")
	}
}

func TestDoGetRequest(t *testing.T) {
	const exampleURL = "https://example.com"
	httpClient := NewHTTPClient()

	res, err := httpClient.DoGetRequest(exampleURL)
	if err != nil {
		t.Fatalf("Expected a nil error, got %s", err)
	}

	if len(res) == 0 {
		t.Fatal("Expected a non-empty array of bytes")
	}
}

func TestDoGetRequestMissingProtocol(t *testing.T) {
	const exampleURL = "://example.com"
	httpClient := NewHTTPClient()

	res, err := httpClient.DoGetRequest(exampleURL)
	if err == nil {
		t.Fatal("Expected an error message, got nil")
	}

	if len(res) != 0 {
		t.Fatal("Expected an empty array of bytes")
	}
}

func TestDoGetRequestInvalidProtocol(t *testing.T) {
	const exampleURL = "XD://example.com"
	httpClient := NewHTTPClient()

	res, err := httpClient.DoGetRequest(exampleURL)
	if err == nil {
		t.Fatal("Expected an error message, got nil")
	}

	if len(res) != 0 {
		t.Fatal("Expected an empty array of bytes")
	}
}

func TestDoGetRequestEmptyURL(t *testing.T) {
	const exampleURL = ""
	httpClient := NewHTTPClient()

	res, err := httpClient.DoGetRequest(exampleURL)
	if err == nil {
		t.Fatal("Expected an error message, got nil")
	}

	if len(res) != 0 {
		t.Fatal("Expected an empty array of bytes")
	}
}
