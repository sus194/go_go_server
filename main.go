package main

import (
	"fmt"
	"go_go_server/data_types"
	"go_go_server/schedulars"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	// Number of worker goroutines
	numWorkers = 10
	Port = ":3000"
)

var(
	id       = 0
    mapMutex = &sync.Mutex{}
)


func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	taskName := r.URL.Path[1:] // Extract task name from URL path
	if _, ok := data_types.TaskMap[taskName]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	request := data_types.Request{
		ResponseWriter: w,
		Request:        r,
		ID:             id,
		ArrivalTime:    time.Now(),
		State:			data_types.ReadyState,
		Priority:		0,
	}
	data_types.RequestMap[taskName] = append(data_types.RequestMap[taskName], request)
	id++
}

func main() {
	mux := http.NewServeMux()
    mux.HandleFunc("/", handleAllRoutes)
    log.Println("Starting server on Port", Port)

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    go func() {
        for range ticker.C {
            schedulars.Schedule()
        }
    }()

    
    if err := http.ListenAndServe(Port, mux); err != nil {
        log.Fatal(err)
    }
}
