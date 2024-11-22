package postfetcher

import (
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"post-storage-service/internal/config"
	adapterMock "post-storage-service/internal/mock/adapter"
	repoMock "post-storage-service/internal/mock/repository"
	"post-storage-service/internal/model"
	"testing"
)

func TestFetchPosts_TotalZero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockPostRepository(ctrl)
	mockPostProvider := adapterMock.NewMockPostProvider(ctrl)
	cfg := config.PostProvider{
		URL:       "test",
		FetchSize: 10,
	}

	totalPosts := 0
	mockPostProvider.EXPECT().GetTotalPosts(gomock.Any()).Return(totalPosts, nil).Times(1)

	postService := NewService(mockRepo, mockPostProvider, cfg)

	ctx := context.Background()
	got := postService.FetchPosts(ctx)

	assert.Equal(t, got, nil)
}

func TestFetchPosts_HundredPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPostRepo := repoMock.NewMockPostRepository(ctrl)
	mockPostProvider := adapterMock.NewMockPostProvider(ctrl)
	cfg := config.PostProvider{
		URL:       "test",
		FetchSize: 10,
	}

	posts := genPosts(10)

	totalPosts := 100
	mockPostProvider.EXPECT().GetTotalPosts(gomock.Any()).Return(totalPosts, nil).Times(1)
	mockPostProvider.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any()).Return(posts, nil).Times(10)

	mockPostRepo.EXPECT().SavePosts(context.Background(), gomock.Eq(posts)).Do(func(ctx context.Context, posts []model.Post) {
		total := len(posts)

		if total != 10 {
			t.Errorf("Expected 10 items, but got %d", total)
		}
	}).Return(nil).Times(10)

	postService := NewService(mockPostRepo, mockPostProvider, cfg)

	ctx := context.Background()
	_ = postService.FetchPosts(ctx)
}

func genPosts(amount int) []model.Post {
	posts := make([]model.Post, amount)
	for i := 1; i <= amount; i++ {
		posts[i-1] = model.Post{
			ID:        uint64(i),
			FirstName: fmt.Sprintf("first name %v", i),
			LastName:  fmt.Sprintf("last name %v", i),
			Email:     fmt.Sprintf("email %v", i),
			Gender:    fmt.Sprintf("gender %v", i),
			IPAddress: fmt.Sprintf("ip address %v", i),
		}
	}

	return posts
}
