package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	source := flag.String("source", "", "Source directory to backup")
	destination := flag.String("destination", "", "Destination directory to save backup")
	flag.Parse()

	if *source == "" || *destination == "" {
		fmt.Println("Please provide source and destination directories")
		os.Exit(1)
	}

	err := filepath.Walk(*source, func(path string, info os.FileInfo, walkErr error) error {

		if walkErr != nil {
			return walkErr
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(*source, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(*destination, relPath)

		err = os.MkdirAll(filepath.Dir(destPath), 0755)
		if err != nil {
			return err
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}

		fmt.Printf("%s backed up to %s\n", path, destPath)
		return nil

	})

	if err != nil {
		fmt.Printf("Error during backup: %v\n", err)
	}
}
