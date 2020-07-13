package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", status)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

