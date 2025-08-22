package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")

	// If the case of an error and the status does not represent an error
	if _, errFound := v.(error); errFound {
		if status < 400 {
			status = http.StatusInternalServerError
		}
	}

	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}
	return err
}
