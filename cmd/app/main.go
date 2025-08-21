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

	listenAddr := os.Getenv("LISTEN_ADDRESS")
	log.Println("API listening on ", listenAddr)
	server := api.NewAPIServer(listenAddr, storage)

	server.Run()
}
