package data_types
import (
	"net/http"
	"time"
	"encoding/json"
)

// Define a custom type for RequestState
type RequestState int
// Define constants for different states using iota
const (
    ReadyState RequestState = iota
	SentState
	ProcessingState
	WaitingState
    CompletedState
    FailedState
)

// Request struct for handling incoming requests
type Request struct {
	// Request details
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ID            int
    ArrivalTime   time.Time
	State  		   RequestState		
}

var(
	RequestMap = make(map[string]Request)
)
