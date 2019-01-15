package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = 8080

// define routes, creating a local ServeMux
// For security reasons, better not to rely on global defaultMux
func getMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/quit", stopHandler)
	return mux
}

//main launches server
func main() {

	fmt.Printf("Starting demo server on port %d\n", port)
	fmt.Printf("Connect to either :\n")
	fmt.Printf("      http://localhost:%d/something\n", port)
	fmt.Printf("      http://localhost:%d/time\n", port)
	fmt.Printf("      http://localhost:%d/quit\n", port)

	log.Fatal(http.ListenAndServe(":8080", getMux()))
}
