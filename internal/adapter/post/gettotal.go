package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	adapterModel "post-storage-service/internal/adapter/post/model"
	"time"
)

// get total posts amount in list
func (a *Adapter) GetTotalPosts(ctx context.Context) (int, error) {
	const op = "adapter.post.GetTotalPosts"
	url := fmt.Sprintf("%s%s?limit=1&offset=0", a.baseURL, getPostsListPath)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//todo: properly handle err
		log.Println(err)
		if errors.Is(err, context.DeadlineExceeded) {
			return 0, fmt.Errorf("%s:%w", op, ErrRequestTimeout)
		}

		return 0, fmt.Errorf("%s:%w", op, ErrInternal)
	}

	defer resp.Body.Close()

	var response adapterModel.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("%s:%w", op, ErrInternal)
	}

	return response.Meta.Total, nil
}
