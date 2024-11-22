package post

import (
	"context"
	"fmt"
	"post-storage-service/internal/model"
	"post-storage-service/internal/repository/pg"
)

type repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *repository {
	return &repository{db: db}
}

// on duplicate just ignores value
func (r *repository) SavePosts(_ context.Context, posts []model.Post) error {
	const op = "repository.pg.SavePosts"
	//data := converter.FromServiceToRepositoryPosts(posts)
	fmt.Println(posts)
	return nil
}
