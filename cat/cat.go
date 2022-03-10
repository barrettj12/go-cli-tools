package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Error: no filename provided as argument. Aborting")
	}

	filenames := os.Args[1:]
	err := readFromFiles(filenames, os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}

func readFromFiles(filenames []string, output io.Writer) error {
	for _, fn := range filenames {
		file, err := os.Open(fn)
		if err != nil {
			return err
		}
		defer file.Close()

		io.Copy(output, file)
	}

	return nil
}
