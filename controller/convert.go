package controller

import (
	"bufio"
	"encoding/json"
	"guepard/model"
	"log"
	"net/http"
	"net/url"
	"os"
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
	resp, err := http.PostForm(postURL, url.Values{"vid": {videoInfo.Vid}, "k": {videoInfo.Kc}})
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
