package post

import (
	"context"
	"fmt"
	"post-storage-service/internal/model"
	"post-storage-service/internal/repository/pg"
	"post-storage-service/internal/repository/post/converter"
)

type repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *repository {
	return &repository{db: db}
}

// on duplicate just ignores value
func (r *repository) SavePosts(ctx context.Context, posts []model.Post) error {
	const op = "repository.pg.SavePosts"
	data := converter.FromServiceToRepositoryPosts(posts)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.NewInsert().
		Model(&data).
		ExcludeColumn("created_at", "updated_at").
		On("CONFLICT(id) DO NOTHING").
		Exec(ctx)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
