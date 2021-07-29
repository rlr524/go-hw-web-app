package middleware

import (
	"fmt"
	"net/http"
)

// WriteToConsole is a test function that writes a message to the console each time a route is hit
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page...")
		next.ServeHTTP(w, r)
	})
}
