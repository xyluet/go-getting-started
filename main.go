package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tag": "1",
			"env": os.Environ(),
		})
	}))
	http.ListenAndServe(":"+port, router)
}
