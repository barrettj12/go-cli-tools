package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

// Constants for testing files
const TESTDIR = "temp/"

var filesToMake = map[string]string{
	"testfile1": "Hello world",
	"testfile2": "testing testing 1 2 3",
}

const nonfile = "this-file-doesn't-exist"

// Setup / teardown for tests
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

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// ACTUAL TESTS START HERE
func TestControl(t *testing.T) {
	// Make buffer to write to
	var testWriter bytes.Buffer
	// Make list of files
	var files []string
	expOutput := ""
	for f, c := range filesToMake {
		files = append(files, TESTDIR+f)
		expOutput += c
	}

	err := readFromFiles(files, &testWriter)
	if err != nil {
		t.Fatalf("Error reading files: %s", err)
	}

	// Check buffer output is as expected
	actualOutput := testWriter.String()
	if actualOutput != expOutput {
		t.Fatalf(`Actual output "%s" doesn't match expected output "%s"`,
			actualOutput, expOutput)
	}
}

func TestFileDoesntExist(t *testing.T) {
	err := readFromFiles(
		[]string{TESTDIR + nonfile},
		nil)
	if err == nil {
		t.Fatalf(`"%s" does not exist, I expected an error`, nonfile)
	}
}

// test files are closed at correct time?
