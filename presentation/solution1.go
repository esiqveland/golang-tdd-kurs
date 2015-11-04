package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// register func
	http.HandleFunc("/message/new", handleMessageNew)
	fmt.Println("serving on http://localhost:7777/message/new")
	// serve func with default server
	log.Fatal(http.ListenAndServe("localhost:7777", http.DefaultServeMux))
}

func handleMessageNew(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if req.Method == "post" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "No such method. Try POST!"}`))
		return
	}

	// START OMIT
//func handleMessageNew(w http.ResponseWriter, req *http.Request) {
	msg := MessageRequest{}
	json.NewDecoder(req.Body).Decode(&msg)
	log.Printf("POST: name=%v", msg.Name)
	resp := MessageResponse{
		Message: fmt.Sprintf("Hello, %v!", msg.Name),
	}
	data, err := json.MarshalIndent(&resp, "", "  ")
	if err != nil {
		log.Fatal("error during marshall: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

type MessageRequest struct {
	Name string `json:"name"`
}
type MessageResponse struct {
	Message string `json:"message"`
}

// END OMIT
