package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", getAPI)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string `json:"msg"`
	}{
		Message: "Hello Shipper",
	}

	json.NewEncoder(w).Encode(response)
}
