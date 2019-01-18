//Package main will launch a demo server for testing various compile options.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// DefPort is the default port used
// Change it by specifying -port on the command line
const DefPort int = 8080

// getMux define routes, creating a local ServeMux
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

	var port int

	// defines flags, bind them to variables, and parse them
	flag.IntVar(&port, "port", DefPort, "define a port")
	flag.IntVar(&port, "p", DefPort, "define a port")
	flag.Parse()

	fmt.Printf("Starting demo server on port %d\n", port)
	fmt.Printf("Connect to either :\n")
	fmt.Printf("      http://localhost:%d/something\n", port)
	fmt.Printf("      http://localhost:%d/time\n", port)
	fmt.Printf("      http://localhost:%d/quit\n", port)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		getMux()))
}
