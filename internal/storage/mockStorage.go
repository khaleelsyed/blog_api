package storage

import (
	"fmt"
	"strings"
	"time"
)

var mockPost Post = Post{
	ID:        2,
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
	var posts = []Post{
		mockPost,
		{
			ID:        3,
			Title:     "My Third Blog Post",
			Content:   "This is the content of my third blog post about COVID.",
			Category:  "Microbiology",
			Tags:      []string{"COVID", "Microbiology", "Biology", "Disease", "Epidemiology"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        4,
			Title:     "My Third Blog Post",
			Content:   "This is the content of my third blog post about magnets.",
			Category:  "Physics",
			Tags:      []string{"Geography", "Space", "Astronomy"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if searchTerm != "" {
		var output = make([]Post, 0)
		for _, post := range posts {
			if checkSearchQueryInPost(post, searchTerm) {
				output = append(output, post)
			}
		}

		return output, nil
	}
	return posts, nil
}

func checkSearchQueryInPost(post Post, searchQuery string) bool {
	query := strings.ToLower(searchQuery)
	if strings.Contains(strings.ToLower(post.Title), query) {
		return true
	} else if strings.Contains(strings.ToLower(post.Content), query) {
		return true
	} else if strings.Contains(strings.ToLower(post.Category), query) {
		return true
	}
	return false
}

func (s MockStorage) Init() error {
	return nil
}

func NewMockStorage() (MockStorage, error) {
	return MockStorage{}, nil
}
