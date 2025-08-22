package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type apiHandlerFunc func(w http.ResponseWriter, r *http.Request) error

type APIServer struct {
	listenAddr string
	storage    Storage
}

func (a APIServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", makeHTTPHandler(a.handlePostsEndpoint))
	r.HandleFunc("/posts/{id}", makeHTTPHandler(a.handlePostsIdEndpoint))

	if err := http.ListenAndServe(a.listenAddr, r); err != nil {

	}
}

func NewAPIServer(listenAddr string, storage Storage) APIServer {
	return APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func makeHTTPHandler(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, err)
		}
	}
}
