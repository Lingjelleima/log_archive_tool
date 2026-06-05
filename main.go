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

func main(){
	if len(os.Args)<2 {
		fmt.Println("Usage: log-archive <log-directory>")
		return
	}
	logDir := os.Args[1]

	//Generate Timestamped Filename
	timestamp := time.Now().Format("20060102")
	archiveName := fmt.Sprintf("logs_archive_%s.tar.gz", timestamp)

	//Create the .tar.gz file
	out,err := os.Create(archiveName)
	if err!=nil {
		fmt.Printf("Error creating archive: %v\n", err)
		return
	}
	defer out.Close()

	//compress
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	//add files to archive
	err=filepath.WalkDir(logDir, func(path string, d os.DirEntry, err error) error {
		if err!=nil{
			return err
		}
		if d.IsDir(){
			return nil
		}

		info,err := d.Info()
		if err!=nil {
			return err
		}
		
		//Create tar header
		header,err := tar.FileInfoHeader(info, "")

		if err!=nil{
			return err
		}
		
		header.Name = filepath.Base(path)

		if err := tw.WriteHeader(header);err!=nil {
			return err
		}
		
		//Copy file content into tar
		file,err := os.Open(path)

		if err!=nil {
			return err
		}
		
		defer file.Close()
		
		_,err=io.Copy(tw, file)
		return err	
	})

	if err!=nil{
		fmt.Printf("Error processing files: %v\n", err)
		return
	}

	//Log to text file
	logEntry := fmt.Sprintf("[%s] Archived %s to %s\n", time.Now().Format(time.RFC3339), logDir, archiveName)
	f,err := os.OpenFile("archive_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err==nil{
		defer f.Close()
		f.WriteString(logEntry)
	}
	fmt.Printf("Successfully created %s\n", archiveName)
}
