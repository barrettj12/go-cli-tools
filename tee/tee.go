package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			break
		} else {
			fmt.Println(text)
			file.WriteString(text)
			file.WriteString("\n")
		}
	}
}
