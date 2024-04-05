package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type responseData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the health check handler")
	resp := responseData{
		Error:   false,
		Message: "Health check successful",
	}
	respByte, err := json.Marshal(&resp)
	if err != nil {
		log.Fatalf("error marshalling json %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respByte)
	if err != nil {
		log.Fatalf("error writing response %s", err)

	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthCheck)

	log.Println("Server is running...")
	http.ListenAndServe("localhost:4000", mux)
}
