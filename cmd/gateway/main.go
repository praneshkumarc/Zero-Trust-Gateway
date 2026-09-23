package main

import (
	"log"
	"net/http"
	"github.com/praneshkumarc/Zero-Trust-Gateway/internal/proxy"
)

func main() {
	handler, err := proxy.New("http://localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Starting Zero Trust Gateway on :8080")
	
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}