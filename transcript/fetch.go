package transcript

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// List ...
func List(videoID string) {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	src, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	build(src)
}

func build(body []byte) {
	data := bytes.Split(body, []byte("\"captions\":"))
	captions := bytes.Split(data[1], []byte(",\"videoDetails"))
	var c Caption
	err := json.Unmarshal(captions[0], &c)
	if err != nil {
		log.Fatal(err)
	}

	extractCaptions(c.PlayerCaptionsTracklistRenderer.CaptionTracks[1])
}

// Transcript ...
type Transcript struct {
	XMLName xml.Name `xml:"transcript"`
	Text    []string `xml:"text"`
}

func extractCaptions(cap CaptionTrack) {
	req, err := http.NewRequest("GET", cap.BaseURL, strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var t Transcript
	err = xml.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		log.Fatal(err)
	}

	for i, text := range t.Text {
		t.Text[i] = html.UnescapeString(text)
	}

	for _, text := range t.Text {
		fmt.Println(text)
	}
}
