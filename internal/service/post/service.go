package post

import "post-storage-service/internal/repository"

type service struct {
	repository repository.PostRepository
}

func NewService(repository repository.PostRepository) *service {
	return &service{repository: repository}
}
