package service

import "context"

type PostService interface {
}

type PostFetcherService interface {
	FetchPosts(ctx context.Context) error
}
