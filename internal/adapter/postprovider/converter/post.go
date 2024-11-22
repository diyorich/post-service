package converter

import (
	"post-storage-service/internal/adapter/postprovider/model"
	svcModel "post-storage-service/internal/model"
)

func FromAdapterToServiceModels(models []model.PostJSON) []svcModel.Post {
	converted := make([]svcModel.Post, len(models))
	for index, value := range models {
		converted[index] = svcModel.Post{
			ID:        value.ID,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
			Gender:    value.Gender,
			IPAddress: value.IPAddress,
		}
	}

	return converted
}
