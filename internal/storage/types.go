package storage

import (
	"fmt"
	"time"
)

type PostContent struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NotFoundErr struct {
	postID int
}

func (e NotFoundErr) Error() string {
	return fmt.Sprintf("Post %d was not found", e.postID)
}
