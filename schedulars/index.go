package schedulars
import(
	"go_go_server/data_types"
)
func Schedule() {
	data_types.Request_Waiting.Dequeue()
	// Implement the Round Robin scheduling algorithm
	// Extract the task handler function from the TaskMap
	// Call the task handler function with the request
	// Update the request state to CompletedState
	// Write the response to the ResponseWriter
}