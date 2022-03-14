package main

import (
	"flag"
	"fmt"
	"os"
)

// COMMAND LINE OPTIONS

// -o flag: set directory where files will be placed
var outputDir = flag.String("o", ".", "directory to download files to")

// -q flag: set max queue size
var bufferSize = flag.Int("q", 100, "max size of URL queue")

// -w flag: set max number of workers
//var numThreads = flag.Int("w", 10, "max number of parallel workers")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("No URL was provided to crawl, hence nothing will be downloaded.")
		fmt.Println("Usage:   crawl [OPTIONS] <url>")
		os.Exit(0)
	}

	runCrawler(flag.Arg(0))
}

// Represents a single web resource to be downloaded
type Task struct {
	source   string // source URL
	destDir  string // destination directory
	filename string // name for downloaded file
}

func runCrawler(url string) {
	// Create new buffered gochannel to use as queue
	taskQueue := make(chan Task, *bufferSize)
	taskQueue <- Task{url, *outputDir, "index.html"}

	for nextTask := range taskQueue {
		go getUrl(nextTask, taskQueue)
	}
}

func getUrl(task Task, taskQueue chan Task) {
	// Make directory to download to
	err := os.MkdirAll(task.destDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create directory %q: %v", task.destDir, err)
		return
	}

	// Make file to write to

	// if external resource encountered: add to `urlQueue`
}
