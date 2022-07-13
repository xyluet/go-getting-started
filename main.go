package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var Version = "development"

func main() {
	port := os.Getenv("XPORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	environs := os.Environ()

	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"version": Version,
			"env":     environs,
			"tag":     1,
		})
	}))
	http.ListenAndServe(":"+port, router)
}
