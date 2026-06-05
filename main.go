package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 1. Validate Command Line Argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: log-archive <log-directory>")
		return
	}
	logDir := os.Args[1]

	// 2. Generate Timestamped Filename
	timestamp := time.Now().Format("20060102_150405")
	archiveName := fmt.Sprintf("logs_archive_%s.tar.gz", timestamp)

	// 3. Create the .tar.gz file
	out, err := os.Create(archiveName)
	if err != nil {
		fmt.Printf("Error creating archive: %v\n", err)
		return
	}
	defer out.Close()

	// Setup compression layers
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// 4. Walk the directory and add files to archive
	filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() { return nil }

		// Create tar header
		header, _ := tar.FileInfoHeader(info, "")
		header.Name = filepath.Base(path)
		tw.WriteHeader(header)

		// Copy file content into tar
		file, _ := os.Open(path)
		defer file.Close()
		io.Copy(tw, file)
		return nil
	})

	// 5. Log the action to a text file
	logEntry := fmt.Sprintf("[%s] Archived %s to %s\n", time.Now().Format(time.RFC3339), logDir, archiveName)
	f, _ := os.OpenFile("archive_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(logEntry)

	fmt.Printf("Successfully created %s\n", archiveName)
}
