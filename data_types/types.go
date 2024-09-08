package data_types
import (
	"net/http"
	"time"
	"go_go_server/data_types/tasks"
	"go_go_server/data_types/queue"
)

// Define a custom type for RequestState
type RequestState int
// Define constants for different states using iota
const (
    ReadyState RequestState = iota
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
	}
	Request_Waiting = &queue.Queue{}
	Request_Processing = &queue.Queue{}
)
