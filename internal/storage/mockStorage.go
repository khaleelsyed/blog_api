package storage

import (
	"fmt"
	"time"
)

var mockPost Post = Post{
	ID:        1,
	Title:     "My Second Blog Post",
	Content:   "This is the content of my second blog post.",
	Category:  "Technology",
	Tags:      []string{"Tech", "Programming"},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type MockStorage struct{}

func (s MockStorage) CreatePost(postContent PostContent) (Post, error) {
	return Post{
		ID:        1,
		Title:     postContent.Title,
		Content:   postContent.Content,
		Category:  postContent.Category,
		Tags:      postContent.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (s MockStorage) UpdatePost(id int, postContent PostContent) (Post, error) {
	return Post{
		ID:        1,
		Title:     postContent.Title,
		Content:   postContent.Content,
		Category:  postContent.Category,
		Tags:      postContent.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (s MockStorage) DeletePost(id int) error {
	if id != 1 {
		return fmt.Errorf("post with id %d not found", id)
	}
	return nil
}

func (s MockStorage) GetPost(id int) (Post, error) {
	return mockPost, nil
}

func (s MockStorage) ListPosts(searchTerm string) ([]Post, error) {
	return []Post{mockPost}, nil
}

func (s MockStorage) Init() error {
	return nil
}

func NewMockStorage() (MockStorage, error) {
	return MockStorage{}, nil
}
