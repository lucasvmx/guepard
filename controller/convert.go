package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"guepard/model"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var musicFolderName = "musicas"

func createMusicsFolder() {
	fail := os.Mkdir(musicFolderName, 0777)
	if fail != nil {
		if os.IsExist(fail) {
			// Just ignore the error
			return
		}

		log.Fatalf("[ERROR] Failed to create folder: %v", fail)
	}
}

func DoConversion(filename string) {

	// Prepare workspace
	createMusicsFolder()

	// Open file in default mode
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalf("[ERROR] Failed to open %v: %v", filename, err)
		return
	}

	// Creates a buffer reader
	reader := bufio.NewReader(file)
	defer file.Close()

	log.Printf("[INFO] Starting download of all musics ...")

	for {
		link, _, fail := reader.ReadLine()
		if fail != nil {
			break
		}

		// Parse string into a URL format
		remoteURL, _ := url.Parse(string(link))

		// Get video information
		videoInfo := GetVideoInformation(remoteURL)

		if videoInfo.Mess != "" {
			log.Printf("[ERROR] %v", videoInfo.Mess)
			continue
		}

		// Request conversion to mp3
		conversionStatus := RequestConversion(videoInfo)

		// Download music
		DownloadMusic(conversionStatus)
	}

	log.Printf("[INFO] All downloads has been finished")
}

func RequestConversion(videoInfo *model.IndexResponse) (status *model.ConvertStatus) {

	log.Printf("[INFO] Requesting video conversion to mp3 format ...")

	postURL := model.GetBaseURL() + model.GetConversionPath()
	payload := fmt.Sprintf("vid=%v&k=%v", url.QueryEscape(videoInfo.Vid), url.QueryEscape(videoInfo.Kc))

	// Build HTTP request
	req, _ := http.NewRequest("POST", postURL, strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", fmt.Sprintf("%d", len(payload)))
	req.Header.Add("User-Agent", "Guepard/0.1")
	req.Header.Add("Host", "yt1s.com")

	// Send HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to request download: %v", err)
		return
	}

	log.Printf("[INFO] Received HTTP code: %v", resp.StatusCode)
	body := readBody(resp)

	// Decode JSON
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Printf("[ERROR] Failed to decode JSON: %v", err)
		return
	}

	return
}
