package transcript

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/SirNoob97/yt-transcripts/client"
)

// Fetcher ...
type Fetcher interface {
	Fetch(videoID, language string) Transcript
	List(videoID string) (map[string]string, error)
}

// Transcript ...
type Transcript struct {
	XMLName xml.Name `xml:"transcript"`
	Text    []string `xml:"text"`
	client  client.Requester
}

// NewTrasncript ...
func NewTrasncript(client client.Requester) Transcript {
	return Transcript{client: client}
}

// List ...
func (t Transcript) List(videoID string) (map[string]string, error) {
	body, err := t.client.DoGetRequest(buildURL(videoID))
	if err != nil {
		return map[string]string{}, err
	}

	captions := getCaptions(body)
	tracks := make(map[string]string, len(captions))
	if len(captions) > 0 {
		for _, track := range captions {
			tracks[track.Name.SimpleText] = track.LanguageCode
		}
	}

	return tracks, nil
}

// Fetch ...
func (t Transcript) Fetch(videoID, language string) Transcript {
	body, err := t.client.DoGetRequest(buildURL(videoID))
	if err != nil {
		log.Fatal(err)
	}

	captions := getCaptions(body)
	for _, track := range captions {
		if language == track.LanguageCode {
			body, err := t.client.DoGetRequest(track.BaseURL)
			if err != nil {
				log.Fatal(err)
			}
			return t.buildTranscript(body)
		}
	}
	return Transcript{}
}

func (t Transcript) buildTranscript(body []byte) Transcript {
	err := xml.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func getCaptions(body []byte) []CaptionTrack {
	data := bytes.Split(body, []byte("\"captions\":"))
	captions := bytes.Split(data[1], []byte(",\"videoDetails"))

	var c *CaptionList
	err := json.Unmarshal(captions[0], &c)
	if err != nil {
		log.Fatal(err)
	}

	return c.PCTR.CaptionTracks
}

func buildURL(videoID string) string {
	if videoID == "" {
		log.Fatal("Video ID is required")
	}
	rep := strings.NewReplacer("\\0026", "&", "\\", "")
	return rep.Replace(fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID))
}
