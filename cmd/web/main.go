/*
Package main provides the application entry point. It spins up a server
and fires off all route handlers.
*/
package main

import (
	"fmt"
	"github.com/rlr524/go-hw-web-app/pkg/handlers"
	"net/http"
)

// In Go, the colon is required when passing the port number in a variable
const portNumber = ":3030"

// main is the application entry point
func main() {
	// Instead of putting our page data into the main() function, here
	// we are still calling the HandleFunc function from the http package but
	// instead we're passing in our page handlers as closures and these
	// functions become our routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/madison", handlers.Madison)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
