package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// register func
	http.HandleFunc("/hello", handleHello)
	fmt.Println("serving on http://localhost:7777/hello")
	// serve func with default server
	log.Fatal(http.ListenAndServe("localhost:7777", http.DefaultServeMux))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello, world!")
}
