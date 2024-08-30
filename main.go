package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"sync"
	"data_types"
	"tasks"
)

const (
	// Number of worker goroutines
	numWorkers = 10
	Port = ":3000"
)

var(
	id            = 0
    mapMutex = &sync.Mutex{}
)

// Worker function to process and forward requests
func worker(id int, requests <-chan Request) {
	for req := range requests {
		// Simulate processing time
		time.Sleep(req.ProcessingTime * time.Millisecond)
		// Process request and send response
		_, err := fmt.Fprintf(req.ResponseWriter, "Worker %d processed request\n", id)
        if err != nil {
            log.Printf("Worker %d: Error writing response: %v\n", id, err)
        }
	}
}

func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	// Lock the map for concurrent access
	mapMutex.Lock()
	defer mapMutex.Unlock()

	// Add request to map with route as key
	route := r.URL.Path
	data_types.RequestMap[route] = Request{
		ResponseWriter: w,
		Request:        r,
		ID:             id,
		ArrivalTime:    time.Now(),
	}
	id++

	// Send a response for demonstration
	fmt.Fprintf(w, "Request for route %s has been added to the map.\n", route)
}

func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	processingTime := 0 * time.Millisecond
	if pt := r.URL.Query().Get("processingTime"); pt != "" {
		if duration, err := time.ParseDuration(pt + "ms"); err == nil {
			processingTime = duration
		}
	}

	requests <- Request{
		ResponseWriter: w,
		Request:        r,
		ID:             id,
		ArrivalTime:    time.Now(),
		ProcessingTime: processingTime,
	}
	id++
}

func main() {
	// Create a channel to hold requests
	requests := make(chan Request, 1000)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, requests)
	}
	// Handler function to enqueue requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleAllRoutes)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Enqueue the request
		var data RequestData

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

		idMutex.Lock()
        currentID := id
        id++
        idMutex.Unlock()

		requests <- Request{
			ResponseWriter: w, 
			Request: r, 
			ID: id, 
			ArrivalTime: time.Now(), 
			ProcessingTime: data.ProcessingTime,
		}
	})

	// Start the server
	log.Println("Starting server on Port", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(err)
	}
}
