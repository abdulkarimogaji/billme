package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abdulkarimogaji/billme/config"
)

type responseData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
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

func StartServer() error {
	server := http.NewServeMux()
	server.HandleFunc("GET /health", healthCheck)

	fmt.Printf("Server is running on port %s", config.AppConfig.PORT)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.AppConfig.PORT), server)
}
