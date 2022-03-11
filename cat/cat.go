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
		err := readFromFile(fn, output)
		if err != nil {
			return err
		}
	}

	return nil
}

func readFromFile(fn string, output io.Writer) (err error) {
	file, err := os.Open(fn)
	// fmt.Printf("Opened file") // %s\n", file.Name())
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(output, file)
	return
}

// func close(file *os.File) {
// 	file.Close()
// 	// fmt.Printf("Closed file") // %s\n", file.Name())
// }
