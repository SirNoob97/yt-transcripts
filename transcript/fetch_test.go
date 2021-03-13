package transcript

import (
	"errors"
	"testing"

	mocks "github.com/SirNoob97/yt-transcripts/mocks/client"
	"github.com/stretchr/testify/mock"
)

func TestNewTranscript(t *testing.T) {
	hc := new(mocks.Requester)
	transcript := NewTrasncript(hc)

	if len(transcript.Text) != 0 {
		t.Fatalf("Expected a Transcript struct with an empty string array")
	}
}

const captionTest = `
{
  "captions":{
    "playerCaptionsTracklistRenderer":{
      "captionTracks":[
        {
          "baseUrl":"BASEURL",
          "name":{
            "simpleText":"SIMPLETEXT"
          },
          "vssId":"VSSID",
          "languageCode":"LANGUAGECODE",
          "kind":"KIND",
          "isTranslatable":true
        }
      ]
    }
  },"videoDetails":"VIDEO DETAILS"
}
`
const textTest = `
<transcript>
	<text>TEXT1</text>
	<text>TEXT2</text>
	<text>TEXT3</text>
</transcript>
`

func TestList(t *testing.T) {
	const videoID = "ID"
	hc := new(mocks.Requester)
	transcript := NewTrasncript(hc)

	hc.On("DoGetRequest", mock.AnythingOfType("string")).Return([]byte(captionTest), nil)

	res, err := transcript.List(videoID)

	if err != nil {
		t.Fatalf("Expected a nil error, got %v", err)
	}

	if len(res) == 0 {
		t.Fatalf("Expected a non-empty array")
	}
}

func TestListFailCase(t *testing.T) {
	const videoID = "ID"
	errorMsg := errors.New("ERROR")
	hc := new(mocks.Requester)
	transcript := NewTrasncript(hc)

	hc.On("DoGetRequest", mock.AnythingOfType("string")).Return([]byte{}, errorMsg)

	res, err := transcript.List(videoID)

	if err == nil {
		t.Fatalf("Expected an error message, got nil")
	}

	if len(res) != 0 {
		t.Fatalf("Expected an empty array, got %s", res)
	}
}

func TestFetch(t *testing.T) {
	const videoID, language, textURL = "ID", "LANGUAGECODE", "BASEURL"
	const videoURL = "https://www.youtube.com/watch?v=ID"
	hc := new(mocks.Requester)
	transcript := NewTrasncript(hc)

	hc.On("DoGetRequest", videoURL).Return([]byte(captionTest), nil)
	hc.On("DoGetRequest", textURL).Return([]byte(textTest), nil)

	res := transcript.Fetch(videoID, language)
	if len(res.Text) == 0 {
		t.Fatal("Expected a non-empty array")
	}
}

func TestFetchFailCase(t *testing.T) {
	const videoID, wronglanguage = "ID", "WRONGLANGUAGE"
	const videoURL = "https://www.youtube.com/watch?v=ID"
	hc := new(mocks.Requester)
	transcript := NewTrasncript(hc)

	hc.On("DoGetRequest", videoURL).Return([]byte(captionTest), nil)

	res := transcript.Fetch(videoID, wronglanguage)
	if len(res.Text) != 0 {
		t.Fatalf("Expected an empty array, got %v", res.Text)
	}
}

func TestBuildTrasncript(t *testing.T) {
	hc := new(mocks.Requester)
	tr := NewTrasncript(hc)
	ct := tr.buildTranscript([]byte(textTest))

	if len(ct.Text) == 0 {
		t.Fatalf("Expected a non-empty array")
	}
}

func TestGetCaptions(t *testing.T) {
	ct := getCaptions([]byte(captionTest))

	if len(ct) == 0 {
		t.Fatalf("Expected a non-empty array")
	}
}

func TestBuildURL(t *testing.T) {
	const videoID = "ID"
	videoURL := "https://www.youtube.com/watch?v=" + videoID

	res := buildURL(videoID)
	if res != videoURL {
		t.Fatalf("Expected \n%s, got \n%s", videoURL, res)
	}
}
