package controller

import (
	"encoding/json"
	"fmt"
	"guepard/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	defaultFormat = "mp3"
)

func readBody(resp *http.Response) []byte {

	// Read entire body at once
	body, fail := ioutil.ReadAll(resp.Body)
	if fail != nil {
		log.Printf("[ERROR] Failed to read body: %v", fail)
	}

	// Close after use
	defer resp.Body.Close()

	return body
}

func GetVideoInformation(youtubeLink *url.URL) (videoInfo *model.IndexResponse) {

	queryParams := &model.IndexQuery{
		Q:  youtubeLink.String(),
		Vt: defaultFormat,
	}

	fullURL := fmt.Sprintf("%v%v", model.GetBaseURL(), model.GetIndexPath())

	log.Printf("[INFO] Getting video information from youtube ...")

	// Send request to get full video information
	resp, _ := http.PostForm(fullURL, url.Values{"q": {queryParams.Q}, "vt": {queryParams.Vt}})

	// Reads HTTP body data
	bodyData := readBody(resp)

	log.Printf("[INFO] Received HTTP code: %v", resp.StatusCode)

	// decode video information
	err := json.Unmarshal(bodyData, &videoInfo)
	if err != nil {
		log.Fatalf("[ERROR] Failed to decode video information from %v: %v", youtubeLink.String(), err)
	}

	return
}
