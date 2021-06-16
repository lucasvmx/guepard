package controller

import (
	"fmt"
	"guepard/model"
	"log"
	"net/http"
	"os"
)

func saveMusic(filename string, musicData []byte) (err error) {

	// Save music into disk
	err = os.WriteFile(filename, musicData, 0777)
	return
}

func buildMusicFilename(videoInfo *model.ConvertStatus) string {
	title := videoInfo.Title

	return fmt.Sprintf("%v/%v.mp3", musicFolderName, title)
}

func DownloadMusic(videoInfo *model.ConvertStatus) {

	log.Printf("[INFO] Downloading music '%v' ...", videoInfo.Title)
	resp, err := http.Get(videoInfo.DownloadLink)
	if err != nil {
		log.Printf("[ERROR] Failed to download music: %v", err)
		return
	}

	// Read HTTP request body
	body := readHTTPFile(resp)

	if body != nil {
		return
	}

	// Save music data into disk
	localFilename := buildMusicFilename(videoInfo)
	err = saveMusic(localFilename, body)
	if err != nil {
		log.Printf("[ERROR] Failed to save music: %v", err)
		return
	}

	log.Printf("[INFO] Music saved: %v", localFilename)

	return
}
