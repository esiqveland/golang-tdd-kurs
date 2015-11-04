package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// register func
	http.HandleFunc("/message", handleMessage)
	fmt.Println("serving on http://localhost:7777/message")
	// serve func with default server
	log.Fatal(http.ListenAndServe("localhost:7777", http.DefaultServeMux))
}

// START OMIT
type MessageResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
}

func handleMessage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp := MessageResponse{
		Message: "Ei lang melding åt deg.",
		Name:    "Åge",
		Age:     31,
	}

	data, err := json.MarshalIndent(&resp, "", "  ") // pointer, WHY??
	if err != nil {
		log.Fatal("error during marshall: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
// END OMIT
