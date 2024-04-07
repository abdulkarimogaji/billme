package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// get pagination
	query := r.URL.Query()
	log.Println(query)

	response := APIResponse{
		Error:   false,
		Message: "Testing",
		Data:    nil,
	}

	respByte, err := json.Marshal(&response)
	if err != nil {
		log.Fatalf("error marshalling json %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respByte)
	if err != nil {
		log.Fatalf("error writing response %s", err)
	}
	// get filters

}
