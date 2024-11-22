package postfetcher

import (
	"post-storage-service/internal/adapter"
	"post-storage-service/internal/config"
	"post-storage-service/internal/repository"
)

type service struct {
	repository   repository.PostRepository
	postProvider adapter.PostProvider
	fetchSize    int
}

func NewService(repository repository.PostRepository, postProvider adapter.PostProvider, cfg config.PostProvider) *service {
	return &service{repository: repository, postProvider: postProvider, fetchSize: cfg.FetchSize}
}
