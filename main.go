package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
  // Folder to scan
	sourceDir := "./logs"
	// Where to save the zip
  archiveDir := "./archive"
    // Filter for files older than this
	daysOld := 3

	// 1. Create archive directory if it doesn't exist
	os.MkdirAll(archiveDir, 0755)

	// 2. Setup the zip file
	zipName := fmt.Sprintf("logs_%s.zip", time.Now().Format("2006-01-02"))
	zipPath := filepath.Join(archiveDir, zipName)
	zipFile, _ := os.Create(zipPath)
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 3. Scan the logs folder
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() { return nil }

		// Check if it's a .log file and older than X days
		if filepath.Ext(path) == ".log" && time.Since(info.ModTime()).Hours() > float64(daysOld*24) {
			fmt.Printf("Archiving: %s\n", info.Name())

			// Add to zip
			f, _ := os.Open(path)
			defer f.Close()
			w, _ := zipWriter.Create(info.Name())
			io.Copy(w, f)

			// 4. Delete the original file
			os.Remove(path)
		}
		return nil
	})

	fmt.Println("Done! Check the archive folder.")
}
