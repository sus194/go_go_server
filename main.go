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


func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	taskName := r.URL.Path[1:] // Extract task name from URL path
	handler, exists := taskMap[taskName]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	request := data_types.Request{
		ResponseWriter: w,
		Request:        r,
		ID:             id,
		ArrivalTime:    time.Now(),
		State:			data_types.ReadyState,
		Priority:		0
	}
	data_types.RequestMap[taskName] = append(data_types.RequestMap[taskName], request)
	id++
	handler(w, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleAllRoutes)

	// Start the server
	log.Println("Starting server on Port", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(err)
	}
}
