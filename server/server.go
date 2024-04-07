package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abdulkarimogaji/billme/config"
	"github.com/abdulkarimogaji/billme/db"
	"github.com/abdulkarimogaji/billme/server/middleware"
)

type responseData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	var resp responseData
	err := db.Storage.Ping()
	if err != nil {
		resp = responseData{
			Error:   true,
			Message: "Failed to ping database",
		}
	} else {
		resp = responseData{
			Error:   false,
			Message: time.Now().String(),
		}
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
	router := http.NewServeMux()
	router.HandleFunc("GET /health", healthCheck)

	fmt.Printf("Server is running on port %s \n", config.AppConfig.PORT)

	stack := middleware.CreateStack(middleware.Logging)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", config.AppConfig.PORT),
		Handler: stack(router),
	}

	return server.ListenAndServe()
}
