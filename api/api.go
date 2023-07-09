package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
   
func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello gorilla/mux!\n"))
}

func main() {
	r := mux.NewRouter()

	// Step 2: Register the route with the handler function
	r.HandleFunc("/", helloMuxHandler)

	// Step 3: Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", r))
}
   