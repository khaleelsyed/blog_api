package api

import "net/http"

type apiError struct {
	Message string `json:"error"`
	Details string `json:"details,omitempty"`
}

func (e apiError) Error() string {
	return e.Message
}

var methodNotAllowed = func(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusMethodNotAllowed, apiError{Message: "method not allowed"})
}
