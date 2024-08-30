package data_types
import (
	"net/http"
	"time"
	"encoding/json"
	"../tasks"
	"data_types/queue"
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
	State  		  RequestState
	Priority	  int		
}
type TaskHandler func(w http.ResponseWriter, r *http.Request)

var(
	RequestMap = make(map[string][]Request)
	TaskMap = map[string]TaskHandler{
		"mouse-click":      tasks.HandleMouseClick,
		"keyboard-input":   tasks.HandleKeyboardInput,
		// Add more tasks here
	}
	Request_Waiting 
)
