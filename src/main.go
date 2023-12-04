// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, World!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello world from api!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	port := 5005
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
