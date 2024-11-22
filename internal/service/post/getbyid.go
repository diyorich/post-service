package post

import (
	"context"
	"post-storage-service/internal/model"
)

func (s *service) GetByID(ctx context.Context, ID uint64) (model.Post, error) {
	return s.repository.GetByID(ctx, ID)
}
