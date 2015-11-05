package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// START OMIT
type MessageResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
}

func main() {
	resp := MessageResponse{
		Message: "Ei lang melding åt deg.",
		Name:    "Åge",
		// Age: 31, // omitempty?!
		Age: 0,
	}

	data, err := json.MarshalIndent(&resp, "", "  ") // pointer, WHY??
	if err != nil {
		log.Fatal("error during marshall: %v", err)
	}
	fmt.Printf("%v", string(data))
}
// END OMIT
