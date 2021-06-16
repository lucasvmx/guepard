package controller

import (
	"encoding/json"
	"fmt"
	"guepard/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	defaultFormat = "mp3"
)

func readHTTPFile(resp *http.Response) []byte {
	buff := make([]byte, 1024)
	fulldata := make([]byte, 1024)
	totalBytes := resp.ContentLength
	var bytesReaded int64 = 0

	if totalBytes <= 0 {
		log.Printf("[CRITICAL] Invalid content length detected (<= 0)")
		return nil
	}

	defer resp.Body.Close()

	// Read file in chunks
	for bytesReaded < totalBytes {
		n, err := resp.Body.Read(buff)
		if err != nil {
			break
		}

		fulldata = append(fulldata, buff...)

		bytesReaded += int64(n)
		log.Printf("Downloaded %v of %v bytes (%v%% complete)", bytesReaded, totalBytes, 100.0*bytesReaded/totalBytes)
	}

	fmt.Printf("\n")

	return fulldata
}

func readBody(resp *http.Response) []byte {

	// Read entire body at once
	body, fail := io.ReadAll(resp.Body)
	if fail != nil {
		log.Printf("[ERROR] Failed to read body: %v", fail)
	}

	// Close after use
	defer resp.Body.Close()

	return body
}

func GetVideoInformation(youtubeLink *url.URL) (videoInfo *model.IndexResponse) {

	fullURL := fmt.Sprintf("%v%v", model.GetBaseURL(), model.GetIndexPath())

	log.Printf("[INFO] Getting video information from youtube ...")

	// Construct request
	body := fmt.Sprintf("q=%v&vt=%v", url.QueryEscape(youtubeLink.String()), defaultFormat)
	req, _ := http.NewRequest("POST", fullURL, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
	req.Header.Add("User-Agent", "Guepard/0.1")
	req.Header.Add("Host", "yt1s.com")

	// Send request to get full video information
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to send HTTP request: %v", err)
		return
	}

	// Reads HTTP body data
	bodyData := readBody(resp)

	log.Printf("[INFO] Received HTTP code: %v", resp.StatusCode)

	// decode video information
	err = json.Unmarshal(bodyData, &videoInfo)
	if err != nil {
		log.Fatalf("[ERROR] Failed to decode video information from %v: %v", youtubeLink.String(), err)
	}

	return
}
