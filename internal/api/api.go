package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiError struct {
	Error string `json:"error"`
}
type apiHandlerFunc func(w http.ResponseWriter, r *http.Request) error

var methodNotAllowed = func(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusMethodNotAllowed, apiError{Error: "method not allowed"})
}

type APIServer struct {
	listenAddr string
	storage    Storage
}

func (a APIServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", makeHTTPHandler(a.handlePostsEndpoint))

	if err := http.ListenAndServe(a.listenAddr, r); err != nil {

	}
}

func NewAPIServer(listenAddr string, storage Storage) APIServer {
	return APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")

	// If the case of an error and the status does not represent an error
	if _, errFound := v.(error); errFound && status < 400 {
		status = http.StatusInternalServerError
	}

	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}
	return err
}

func makeHTTPHandler(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, err)
		}
	}
}
