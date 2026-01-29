package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler)

	log.Print("Listening...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
