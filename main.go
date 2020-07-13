package main

import (
	"log"
	"net/http"
)

func main() {
	InitLogs()

	http.HandleFunc("/status", status)
	http.HandleFunc("/consentUpdate", handleConsentUpdate)
	http.HandleFunc("/dataflowUpdate", handleDataflowUpdate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
