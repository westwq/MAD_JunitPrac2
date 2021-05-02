package main

import (
	"fmt"
	"io"
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
	absPath, err = findFileAbsPath("../app", "ExampleInstrumentedTest.java")
	if err != nil {
		log.Fatal(err)
	}
	realPathAnd := absPath[:len(absPath)-28]
	fmt.Println(realPathAnd)

	fromAnd, err := os.Open("../testDirectory/MainActivityTest.java")
	if err != nil {
		log.Fatal(err)
	}
	defer fromAnd.Close()

	toAnd, err := os.OpenFile(realPathAnd+"MainActivityTest.java", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer toAnd.Close()

	_, err = io.Copy(toAnd, fromAnd)
	if err != nil {
		log.Fatal(err)
	}

	absPath, err = findFileAbsPath("../app", "ExampleUnitTest.java")
	if err != nil {
		log.Fatal(err)
	}
	realPath := absPath[:len(absPath)-20]
	fmt.Println(realPath)

	from, err := os.Open("../testDirectory/Practice2Test.java")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(realPath+"UserTest.java", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

	from, err = os.Open("../testDirectory/MainActivityTest.java")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err = os.OpenFile(realPath+"MainActivityTest.java", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
