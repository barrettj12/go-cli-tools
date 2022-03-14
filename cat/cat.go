package main

import (
	"flag"
	"io"
	"log"
	"os"
)

// -o flag: direct output to file
var output = flag.String("o", "", "file to write output to")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) <= 1 {
		log.Fatal("Error: expected filename as argument but none was provided.\n" +
			"Usage:  cat [-o <outputfile>] <file1> [<file2> ...]\n" +
			"Aborting.")
	}

	filenames := args[1:]
	writeout := os.Stdout
	if *output != "" {
		var err error
		writeout, err = os.Create(*output)
		if err != nil {
			log.Fatal("Failed to create output file:", err)
		}
	}

	err := readFromFiles(filenames, writeout)
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

// Separate function - defer ensures that files close when done
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
