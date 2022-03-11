package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

const TESTDIR = "temp/"

var filesToMake = map[string]string{
	"testfile1": "Hello world",
	"testfile2": "testing testing 1 2 3",
}

const nonfile = "this-file-doesn't-exist"

func setup() {
	os.Mkdir(TESTDIR, os.ModePerm)

	// Make testfiles
	for fn, contents := range filesToMake {
		err := makeFile(fn, contents)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Ensure `nonfile` doesn't exist
	err := os.Remove(TESTDIR + nonfile)
	if !(err != nil || errors.Is(err, fs.ErrNotExist)) {
		// fmt.Println(err)
		log.Fatal(err)
	}
}

func makeFile(fn string, contents string) (err error) {
	file, err := os.Create(TESTDIR + fn)
	if err != nil {
		fmt.Println("couldn't create", fn)
		return
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	return
}

func teardown() {
	// Remove temporary files and testdir
	err := os.RemoveAll(TESTDIR)
	if err != nil {
		log.Fatal(err)
	}
}

func TestControl(t *testing.T) {
	// testWriter =
}

// Testcontrol

// testfiledoensntexist

// test files are closed at correct time

func main() {
	teardown()
}
