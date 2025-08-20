package api

import "github.com/khaleelsyed/blog_api/internal/storage"

type Storage interface {
	CreatePost(postContent storage.PostContent) (storage.Post, error)
	UpdatePost(id int, postContent storage.PostContent) (storage.Post, error)
	DeletePost(id int) error
	GetPost(id int) (storage.Post, error)
	ListPosts(searchTerm string) ([]storage.Post, error)
}
