package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khaleelsyed/blog_api/internal/storage"
)

func (a APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	var postCreate storage.PostContent

	if err := json.NewDecoder(r.Body).Decode(&postCreate); err != nil {
		log.Println(err)
		return WriteJSON(w, http.StatusBadRequest, apiError{Message: "invalid request body", Details: err.Error()})
	}

	post, err := a.storage.CreatePost(postCreate)
	if err != nil {
		return apiError{Message: "there was an issue loading to the database", Details: err.Error()}
	}

	return WriteJSON(w, http.StatusCreated, post)
}

func (a APIServer) handleListPosts(w http.ResponseWriter, r *http.Request) error {
	queryString := r.URL.Query().Get("term")
	posts, err := a.storage.ListPosts(queryString)
	if err != nil {
		log.Println(err)
		return WriteJSON(w, http.StatusBadRequest, apiError{Message: "invalid query parameters", Details: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, posts)
}

func (a APIServer) handleGetPost(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, apiError{Message: "enter a valid ID"})
	}

	post, err := a.storage.GetPost(id)
	if err != nil {
		if errors.As(err, &storage.NotFoundErr{}) {
			return WriteJSON(w, http.StatusNotFound, apiError{Message: err.Error()})
		}
		return WriteJSON(w, http.StatusInternalServerError, apiError{Message: "failed to get post", Details: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, post)

}

func (a APIServer) handleUpdatePost(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, apiError{Message: "enter a valid ID"})
	}

	var postUpdate storage.PostContent

	if err = json.NewDecoder(r.Body).Decode(&postUpdate); err != nil {
		log.Println(err)
		return WriteJSON(w, http.StatusBadRequest, apiError{Message: "invalid request body", Details: err.Error()})
	}

	post, err := a.storage.UpdatePost(id, postUpdate)
	if err != nil {
		if errors.As(err, &storage.NotFoundErr{}) {
			return WriteJSON(w, http.StatusNotFound, apiError{Message: err.Error()})
		}
		return WriteJSON(w, http.StatusInternalServerError, apiError{Message: "failed to update post", Details: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, post)
}

func (a APIServer) handleDeletePost(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, apiError{Message: "enter a valid ID"})
	}

	if err = a.storage.DeletePost(id); err != nil {
		if err != nil {
			if errors.As(err, &storage.NotFoundErr{}) {
				return WriteJSON(w, http.StatusNotFound, apiError{Message: err.Error()})
			}
			return WriteJSON(w, http.StatusInternalServerError, apiError{Message: "failed to update post", Details: err.Error()})
		}
	}

	return WriteJSON(w, http.StatusNoContent, nil)
}

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

func (a APIServer) handlePostsIdEndpoint(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return a.handleGetPost(w, r)
	case http.MethodPut:
		return a.handleUpdatePost(w, r)
	case http.MethodDelete:
		return a.handleDeletePost(w, r)
	default:
		return methodNotAllowed(w)
	}

}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, err
	}
	return id, err
}
