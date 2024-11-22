package post

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/diyorich/post-api/pkg"
	"github.com/pkg/errors"
	"post-storage-service/internal/model"
	repoErr "post-storage-service/internal/repository"
	"post-storage-service/internal/repository/pg"
	"post-storage-service/internal/repository/post/converter"
	repoModel "post-storage-service/internal/repository/post/model"
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
		On("CONFLICT (id) DO UPDATE " +
			"SET first_name = EXCLUDED.first_name, " +
			"last_name=EXCLUDED.last_name, " +
			"gender = EXCLUDED.gender, " +
			"ip_address=EXCLUDED.ip_address, " +
			"email = EXCLUDED.email, " +
			"updated_at = now()").
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

func (r *repository) GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error) {
	const op = "repository.post.GetList"

	var posts []repoModel.Post
	totalRows, err := r.db.NewSelect().
		Model(&posts).
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		ScanAndCount(ctx)

	if err != nil {
		return []model.Post{}, fmt.Errorf("%s:%w", op, err)
	}

	pagination.Total = totalRows

	return converter.FromRepositoryToServicePosts(posts), nil
}

func (r *repository) GetByID(ctx context.Context, ID uint64) (model.Post, error) {
	const op = "repository.post.GetByID"

	var post repoModel.Post
	err := r.db.NewSelect().
		Model(&post).
		Where("id = ?", ID).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Post{}, fmt.Errorf("%s:%w", op, repoErr.ErrPostNotFound)
		}

		return model.Post{}, fmt.Errorf("%s:%w", op, repoErr.ErrInternal)
	}

	return converter.FromRepositoryToServicePost(post), nil
}
