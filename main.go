package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":3030"

// Home is the / page handler. All web handler functions need to take in, as params,
// the ResponseWriter method and a pointer to the Request method
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is the home page.")
	if err != nil {
		fmt.Println(err)
	}
}

// About is the /about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(10, 10)
	_, err := fmt.Fprintf(w, fmt.Sprintf("This is the about page. The sum of the numbers is %d", sum))
	if err != nil {
		fmt.Println(err)
	}
}

// addValues adds two integers and returns the sum. Remember in Go, a lowercase func name or var name is private
func addValues(x, y int) int {
	sum := x + y
	return sum
}

// Divide is the /divide page handler
func Divide (w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	_, e := fmt.Fprintf(w, fmt.Sprintf( "%f divided by %f is %f", 100.0, 0.0, f))
	if e != nil {
		fmt.Println(e)
	}
}

// divideValues divides two integers and returns the quotient
func divideValues (x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	quotient := x / y
	return quotient, nil
}

// main is the application entry point
func main() {
	// Instead of putting our page data into the main() function, here
	// we are still calling the HandleFunc function from the http package but
	// instead we're passing in our Home handler as a closure
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	
	
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello, Nicole!")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	//})
	
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
	}
