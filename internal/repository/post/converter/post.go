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
