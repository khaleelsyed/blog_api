package api

import "net/http"

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
	return WriteJSON(w, http.StatusNotImplemented, "create post not implemented")
}

func (a APIServer) handleListPosts(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusNotImplemented, "list posts not implemented")
}
