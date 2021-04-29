package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var absPath string // global variable to store absolute path of filename.

func findFileAbsPath(basepath, filename string) (string, error) {
	// This function finds the absolute path of specified filename, returns the absolute path filename and error.
	// The function takes in a base path to search from, filepath.Walk searches file recursively.
	fileErr := filepath.Walk(basepath, func(path string, info os.FileInfo, err error) error {
		// path is the absolute path.
		if err != nil {
			return err
		}
		if info.Name() == filename {
			// if the filename is found return the absolute path of the file.
			absPath = path
		}
		return nil
	})
	if fileErr != nil {
		return "", fileErr
	}
	return absPath, nil
}

func main() {
	// usage example.
	var err error
	absPath, err = findFileAbsPath("..\app", "ExampleUnitTest.java")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(absPath)
}
