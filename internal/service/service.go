package service

import (
	"context"
	"github.com/diyorich/post-api/pkg"
	"post-storage-service/internal/model"
)

type PostService interface {
	GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error)
	GetByID(ctx context.Context, ID uint64) (model.Post, error)
}

type PostFetcherService interface {
	FetchPosts(ctx context.Context) error
}
