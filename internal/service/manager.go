package service

import (
	"post-storage-service/internal/adapter/post"
	"post-storage-service/internal/config"
	"post-storage-service/internal/repository/pg"
	repository "post-storage-service/internal/repository/post"
	"post-storage-service/internal/service/postfetcher"
)

type Manager struct {
	PostService        PostService
	PostFetcherService PostFetcherService
}

func NewManager(db *pg.DB, cfg config.PostProvider) *Manager {
	return &Manager{
		PostFetcherService: postfetcher.NewService(repository.NewRepository(db), post.NewPostAdapter(cfg.URL), cfg),
	}
}
