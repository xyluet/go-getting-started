package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("XPORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tag": "10",
			"env": os.Environ(),
		})
	}))
	http.ListenAndServe(":"+port, router)
}
