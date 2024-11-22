package adapter

import (
	"context"
	"post-storage-service/internal/model"
)

type PostProvider interface {
	Fetch(ctx context.Context, limit int, offset int) ([]model.Post, error)
	GetTotalPosts(ctx context.Context) (int, error)
}
