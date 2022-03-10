package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	files := os.Args[1:]
	readers := make([]io.Reader, len(files))

	for i, file := range files {
		readers[i], _ = os.Open(file)
	}

	mr := io.MultiReader(readers...)

	// Read off `mr`
	buf := make([]byte, 8)
	for {
		n, err := mr.Read(buf)
		fmt.Printf("%s", buf[0:n])
		if err != nil {
			break
		}
	}
}
