package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var port = 8080

// defaultHandler print a message with the path information
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var p = r.URL.Path
	fmt.Fprintf(w, "The path was : %s", p)
}

// timeHandler give current server time
func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

// stopHandler triggers a server stop 2 seconds later
func stopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nThe server will stop in 2 seconds")
	// Using a goroutine, to allow response time to be sent ...
	go stopLater(2 * time.Second)
}

// stopLater waits for duration before stoppping everything
func stopLater(d time.Duration) {
	fmt.Printf("Server will stop in %v  \n", d)
	time.Sleep(d)
	fmt.Println("Server stopped !")
	os.Exit(0)
}

// define routes, creting a local ServeMux
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
