package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	// Number of worker goroutines
	numWorkers = 10
)

// Request struct for handling incoming requests
type Request struct {
	// You can add more fields here if needed
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// Worker function to process requests
func worker(id int, requests <-chan Request) {
	for req := range requests {
		// Simulate processing time
		time.Sleep(500 * time.Millisecond)

		// Process request and send response
		fmt.Fprintf(req.ResponseWriter, "Worker %d processed request\n", id)
	}
}

func main() {
	// Create a channel to hold requests
	requests := make(chan Request, 100)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, requests)
	}

	// Handler function to enqueue requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Enqueue the request
		requests <- Request{ResponseWriter: w, Request: r}
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
