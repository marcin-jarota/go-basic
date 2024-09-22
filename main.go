package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	http.HandleFunc("/", handleRoot)
	log.Println("Server starting on port 3002...")
	if err := http.ListenAndServe(":3002", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
		Version: "Go " + getGoVersion(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getGoVersion() string {
	return "1.20+" // This is a placeholder. In a real app, you might use runtime.Version()
}
