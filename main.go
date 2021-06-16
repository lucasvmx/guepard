package main

import (
	"fmt"
	"guepard/controller"
	"log"
	"os"
)

func main() {
	var filename string

	// Require command line arguments
	if len(os.Args) == 2 {
		filename = os.Args[1]

		if filename == "-h" || filename == "--help" || filename == "/?" {
			fmt.Printf("\nUsage: %v <filename\n", os.Args[0])
			return
		}

	} else {
		log.Printf("[INFO] Filename not informed. Using 'musicas.txt'")
		filename = "musicas.txt"
	}

	// Read the file line by line and download videos
	controller.DoConversion(filename)
}
