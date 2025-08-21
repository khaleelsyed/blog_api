package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/khaleelsyed/blog_api/internal/storage"
)

func (a APIServer) handlePostsEndpoint(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		return a.handleCreatePost(w, r)
	case http.MethodGet:
		return a.handleListPosts(w, r)
	default:
		return methodNotAllowed(w)
	}
}

func (a APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	var postCreate storage.PostContent

	if err := json.NewDecoder(r.Body).Decode(&postCreate); err != nil {
		log.Println(err)
		return WriteJSON(w, http.StatusBadRequest, apiError{Error: "invalid request body", Details: err.Error()})
	}

	post, err := a.storage.CreatePost(postCreate)
	if err != nil {
		return errors.New("there was an issue loading to the database")
	}

	return WriteJSON(w, http.StatusCreated, post)
}

func (a APIServer) handleListPosts(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusNotImplemented, apiError{Error: "list posts not implemented"})
}
