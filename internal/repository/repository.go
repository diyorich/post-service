package repository

import (
	"context"
	"post-storage-service/internal/model"
)

type PostRepository interface {
	SavePosts(ctx context.Context, posts []model.Post) error
}
