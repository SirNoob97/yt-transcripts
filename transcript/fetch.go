package transcript

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	captions := bytes.Split(data[1], []byte(", \"videoDetails"))
	fmt.Println(string(captions[0]))
	var c Caption
	err := json.Unmarshal(captions[0], &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
