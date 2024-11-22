package postfetcher

import (
	"context"
	"post-storage-service/internal/config"
	"post-storage-service/internal/model"
	"post-storage-service/internal/repository"
)

type PostProvider interface {
	Fetch(ctx context.Context, limit int, offset int) ([]model.Post, error)
	GetTotalPosts(ctx context.Context) (int, error)
}

type service struct {
	repository   repository.PostRepository
	postProvider PostProvider
	fetchSize    int
}

func NewService(repository repository.PostRepository, postProvider PostProvider, cfg config.PostProvider) *service {
	return &service{repository: repository, postProvider: postProvider, fetchSize: cfg.FetchSize}
}
