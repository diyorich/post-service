package converter

import (
	"post-storage-service/internal/model"
	repoModel "post-storage-service/internal/repository/post/model"
)

func FromServiceToRepositoryPosts(posts []model.Post) []repoModel.Post {
	conv := make([]repoModel.Post, len(posts))
	for index, value := range posts {
		if !value.IsValidGender() {
			value.Gender = model.NonBinaryGender
		}

		conv[index] = repoModel.Post{
			ID:        value.ID,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
			Gender:    value.Gender,
			IPAddress: value.IPAddress,
		}
	}

	return conv
}

func FromRepositoryToServicePosts(posts []repoModel.Post) []model.Post {
	conv := make([]model.Post, len(posts))
	for index, value := range posts {
		conv[index] = model.Post{
			ID:        value.ID,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
			Gender:    value.Gender,
			IPAddress: value.IPAddress,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
	}

	return conv
}

func FromRepositoryToServicePost(post repoModel.Post) model.Post {
	return model.Post{
		ID:        post.ID,
		FirstName: post.FirstName,
		LastName:  post.LastName,
		Email:     post.Email,
		Gender:    post.Gender,
		IPAddress: post.IPAddress,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
