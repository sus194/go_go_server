package main

import (
	"io"
	"log"
	"net/http"
)

const (
	// Number of worker goroutines
	numWorkers = 10
	// Target API server URL
	apiServerURL = "http://api.example.com"
	Port string = ":3000"
)

// Request struct for handling incoming requests
type Request struct {
	// Request details
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// Worker function to process and forward requests
func worker(id int, requests <-chan Request) {
	client := &http.Client{}
	for req := range requests {
		// Create a new request to forward to the API server
		apiReq, err := http.NewRequest(req.Request.Method, apiServerURL+req.Request.URL.Path, req.Request.Body)
		if err != nil {
			log.Printf("Worker %d: Error creating request: %v", id, err)
			req.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			continue
		}

		// Copy headers
		for name, values := range req.Request.Header {
			for _, value := range values {
				apiReq.Header.Add(name, value)
			}
		}

		// Forward request
		resp, err := client.Do(apiReq)
		if err != nil {
			log.Printf("Worker %d: Error forwarding request: %v", id, err)
			req.ResponseWriter.WriteHeader(http.StatusBadGateway)
			continue
		}
		defer resp.Body.Close()

		// Copy the API server response
		req.ResponseWriter.WriteHeader(resp.StatusCode)
		io.Copy(req.ResponseWriter, resp.Body)
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
	log.Println("Starting server on Port", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(err)
	}
}
