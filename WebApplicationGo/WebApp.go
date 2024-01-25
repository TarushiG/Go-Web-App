package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new ServeMux instance
	mux := http.NewServeMux()

	// Define your route and handler function
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Basic Web Application")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// Start the HTTP server with the ServeMux as the handler
	http.ListenAndServe(":3000", mux)
}
