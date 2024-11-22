package main

import (
	"context"
	"log"
	"post-storage-service/internal/config"
	"post-storage-service/internal/repository/pg"
	"post-storage-service/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	//connect to db
	db, err := pg.Dial(*cfg)
	if err != nil {
		return err
	}

	svcManager := service.NewManager(db, cfg.PostProvider)

	err = svcManager.PostFetcherService.FetchPosts(context.Background())
	if err != nil {
		return err
	}

	return nil
}
