package repository

import (
	"context"
	"github.com/diyorich/post-api/pkg"
	"post-storage-service/internal/model"
)

type PostRepository interface {
	SavePosts(ctx context.Context, posts []model.Post) error
	GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error)
	GetByID(ctx context.Context, ID uint64) (model.Post, error)
}
