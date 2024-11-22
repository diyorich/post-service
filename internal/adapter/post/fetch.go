package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"post-storage-service/internal/adapter/post/converter"
	adapterModel "post-storage-service/internal/adapter/post/model"
	"post-storage-service/internal/model"
	"time"
)

func (a *Adapter) Fetch(ctx context.Context, limit int, offset int) ([]model.Post, error) {
	const op = "adapter.post.Fetch"
	url := fmt.Sprintf("%s%s?limit=%v&offset=%v", a.baseURL, getPostsListPath, limit, offset)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []model.Post{}, fmt.Errorf("%s: %w", op, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//todo: properly handle err
		log.Println(err)
		if errors.Is(err, context.DeadlineExceeded) {
			return []model.Post{}, fmt.Errorf("%s:%w", op, ErrRequestTimeout)
		}

		return []model.Post{}, fmt.Errorf("%s:%w", op, ErrInternal)
	}

	defer resp.Body.Close()

	var response adapterModel.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return []model.Post{}, fmt.Errorf("%s:%w", op, ErrInternal)
	}

	return converter.FromAdapterToServiceModels(response.Posts), nil
}
