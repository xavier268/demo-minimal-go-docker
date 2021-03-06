// Package main defines a set of demo handlers it uses
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

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
