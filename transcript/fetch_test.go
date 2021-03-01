package transcript

import (
	"testing"

	mocks "github.com/SirNoob97/yt-transcripts/mocks/client"
)

func TestNewTranscript(t *testing.T) {
	hc := new(mocks.Requester)
	transcript := NewTrasncript(hc)

	if len(transcript.Text) != 0 {
		t.Fatalf("Expected a Transcript struct with an empty string array")
	}
}

const jsonTest = `
{
  "captions" : {
    "playerCaptionsTracklistRenderer" : {
      "captionTracks" : [
        {
          "baseUrl" : "BASEURL",
          "name" : {
            "simpleText" : "SIMPLETEXT"
          },
          "vssId" : "VSSID",
          "languageCode" : "LANGUAGECODE",
          "kind" : "KIND",
          "isTranslatable" : "ISTRANSLATABLE"
        }
      ]
    }
  }
}
`
