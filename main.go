package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response struct for JSON response
type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/msg", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "My server is working"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
