// USAGE: wget <url> <filename>
package main

import (
	"net/http"
	"os"
)

func main() {
	file, _ := os.Create(os.Args[2])
	client := &http.Client{}
	resp, _ := client.Get(os.Args[1])
	webreader := resp.Body
	buf := make([]byte, 8)

	for {
		n, err := webreader.Read(buf)
		file.Write(buf[:n])
		if err != nil {
			break
		}
	}
}
