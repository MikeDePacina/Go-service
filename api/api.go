package main

import (
	"fmt"
	"log"
	"net/http"
)
   
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello World")
}

func main() {
	

	// Step 2: Register the route with the handler function
	http.HandleFunc("/", helloHandler)

	// Step 3: Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
   