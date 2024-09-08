package tasks

import (
	"net/http"
)

func HandleMouseClick(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Mouse clicked"))
}

func HandleKeyboardInput(w http.ResponseWriter, r *http.Request){
	
}