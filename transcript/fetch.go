package transcript

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// ListTranscripts ...
func ListTranscripts(videoID string, client *http.Client) ([]string, error) {
	captions := getCaptions(videoID, client)
	tracks := make([]string, 0, len(captions))

	if len(captions) > 0 {
		tracks = append(tracks, fmt.Sprintf("Available Transcripts of %s", videoID))
		for i, track := range captions {
			tracks = append(tracks, fmt.Sprintf("Transcript #%d - %s(%s)", i+1, track.Name.SimpleText, track.LanguageCode))
		}
	}

	return tracks, nil
}

// FetchTranscript ...
func FetchTranscript(videoID, language string, client *http.Client) Transcript {
	captions := getCaptions(videoID, client)

	if language == "" {
		language = captions[0].LanguageCode
	}

	for _, track := range captions {
		if language == track.LanguageCode {
			body, err := getRequest(track.BaseURL, client)
			if err != nil {
				log.Fatal(err)
			}
			return buildTranscript(body)
		}
	}
	return Transcript{}
}

// Transcript ...
type Transcript struct {
	XMLName xml.Name `xml:"transcript"`
	Text    []string `xml:"text"`
}

func buildTranscript(body []byte) Transcript {
	var t Transcript
	err := xml.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func getCaptions(videoID string, client *http.Client) []CaptionTrack {
	body, err := getRequest(buildURL(videoID), client)
	if err != nil {
		log.Fatal(err)
	}

	data := bytes.Split(body, []byte("\"captions\":"))
	captions := bytes.Split(data[1], []byte(",\"videoDetails"))

	var c *CaptionList
	err = json.Unmarshal(captions[0], &c)
	if err != nil {
		log.Fatal(err)
	}

	return c.PCTR.CaptionTracks
}

func getRequest(url string, client *http.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func buildURL(videoID string) string {
	if videoID == "" {
		fmt.Println("Video ID is required")
		os.Exit(1)
	}
	rep := strings.NewReplacer("\\0026", "&", "\\", "")
	return rep.Replace(fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID))
}
