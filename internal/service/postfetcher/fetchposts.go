package postfetcher

import (
	"context"
	"fmt"
)

// FetchPosts fetches all posts from post Provider
func (s *service) FetchPosts(ctx context.Context) error {
	const op = "service.postfetcher.FetchPosts"
	totalPosts, err := s.postProvider.GetTotalPosts(ctx)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	if totalPosts == 0 {
		return nil
	}

	var fetchedAmount, offset int

	var fetchErrs []error
	var saveErrs []error
	for fetchedAmount < totalPosts {
		fetchedAmount += s.fetchSize
		posts, err := s.postProvider.Fetch(ctx, s.fetchSize, offset)
		if err != nil {
			fetchErrs = append(fetchErrs, err)
			continue
		}

		//todo: move to queue
		sErr := s.repository.SavePosts(ctx, posts)
		if sErr != nil {
			saveErrs = append(saveErrs, sErr)
			continue
		}

		offset += s.fetchSize
	}

	if fetchErrs != nil || saveErrs != nil {
		return &FetchError{
			FetchErrs: fetchErrs,
			SaveErrs:  saveErrs,
		}
	}

	return nil
}
