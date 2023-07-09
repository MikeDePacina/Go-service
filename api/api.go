package main

import (
	"log"
	"net/http"
)
   func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I love Melissa!\n"))
   }
   func main() {
	http.HandleFunc("/", helloGoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
   }
   