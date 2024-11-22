package service

import (
	"post-storage-service/internal/adapter/post"
	"post-storage-service/internal/config"
	"post-storage-service/internal/repository/pg"
	repository "post-storage-service/internal/repository/post"
	post2 "post-storage-service/internal/service/post"
	"post-storage-service/internal/service/postfetcher"
)

type Manager struct {
	PostService        PostService
	PostFetcherService PostFetcherService
}

func NewManager(db *pg.DB, cfg config.PostProvider) *Manager {
	postRepository := repository.NewRepository(db)

	return &Manager{
		PostFetcherService: postfetcher.NewService(postRepository, post.NewPostAdapter(cfg.URL), cfg),
		PostService:        post2.NewService(postRepository),
	}
}
