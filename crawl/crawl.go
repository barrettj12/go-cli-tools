package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// COMMAND LINE OPTIONS

// -o flag: set directory where files will be placed
var outputDir = flag.String("o", ".", "directory to download files to")

// -q flag: set max queue size
var bufferSize = flag.Int("q", 100, "max size of URL queue")

// -w flag: set max number of workers
var numWorkers = flag.Int("w", 10, "max number of parallel workers")

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

	// Create channel of "tokens"
	// The main thread must use a token to create a new goroutine
	// When the goroutine terminates, it returns the token
	tokens := make(chan int, *numWorkers)
	for i := 0; i < *numWorkers; i++ {
		tokens <- 0
	}

	// Create error channel
	errorChan := make(chan error, *numWorkers)

	for {
		// Process errors first
		for len(errorChan) > 0 {
			fmt.Fprintf(os.Stderr, "Error: %v", <-errorChan)
		}

		if len(taskQueue) > 0 {
			// Get next job
			task := <-taskQueue
			// Wait for another token
			<-tokens
			// Start new goroutine
			go getUrl(task, taskQueue, tokens, errorChan)
		} else {
			// No more jobs
			if len(tokens) == *numWorkers {
				// No goroutines running, hence all tasks are done
				os.Exit(0)
			}
			// Else - need to wait as a currently running worker
			//   could create a new job
		}
	}
}

func getUrl(task Task, taskQueue chan Task, tokens chan int, errorChan chan error) {
	// Make sure to return token at end
	defer func() { tokens <- 0 }()

	// Make directory to download to
	err := os.MkdirAll(task.destDir, os.ModePerm)
	if err != nil {
		errorChan <- fmt.Errorf("Could not create directory %q: %w", task.destDir, err)
		return
	}

	// Make file to write to
	file, err := os.Create(filepath.Join(task.destDir, task.filename))
	if err != nil {
		errorChan <- fmt.Errorf("Could not create file %q in dir %q: %w", task.filename, task.destDir, err)
		return
	}

	// if external resource encountered: add to `taskQueue`
}
