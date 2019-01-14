package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	var p = r.URL.Path
	fmt.Fprintf(w, "Hi there, you typed %s", p)
	if p == "/quit" {
		// This will never appear. Why ?
		fmt.Fprintf(w, "\nStopping the server in 2 seconds")
		// Using a goroutine, to allow response time to be sent ...
		go stopLater(2 * time.Second)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

// stopLater waits for duration before stoppping everything
func stopLater(d time.Duration) {
	fmt.Printf("Server will stop in %v  \n", d)
	time.Sleep(d)
	fmt.Println("Server stopped !")
	os.Exit(43)
}

//
func main() {

	fmt.Println("Starting demo server on port 8080")
	fmt.Println("Connect to either :")
	fmt.Println("      http://localhost:8080/something")
	fmt.Println("      http://localhost:8080/time")
	fmt.Println("      http://localhost:8080/quit")

	http.HandleFunc("/", testHandler)
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
