package main

import (
	"log"
	"os"

	"github.com/khaleelsyed/blog_api/internal/api"
	"github.com/khaleelsyed/blog_api/internal/storage"
)

func main() {
	storage, err := storage.NewMockStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err = storage.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(os.Getenv("LISTEN_ADDRESS"), storage)

	server.Run()
}
