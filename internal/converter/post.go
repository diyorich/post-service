package converter

import (
	"post-storage-service/internal/model"
)

func FromServiceToPostsJSON(posts []model.Post) []model.PostJSON {
	conv := make([]model.PostJSON, len(posts))
	for index, value := range posts {
		conv[index] = model.PostJSON{
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

func FromServiceToPostJSON(post model.Post) model.PostJSON {
	return model.PostJSON{
		ID:        post.ID,
		FirstName: post.FirstName,
		LastName:  post.LastName,
		Email:     post.Email,
		Gender:    post.Gender,
		IPAddress: post.IPAddress,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
