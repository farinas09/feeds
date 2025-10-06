package search

import (
	"context"

	"github.com/farinas09/feeds-api/models"
)

type SearchRepository interface {
	Close()
	IndexFeed(ctx context.Context, feed *models.Feed) error
	SearchFeeds(ctx context.Context, query string) ([]*models.Feed, error)
}

var searchRepository SearchRepository

func SetSearchRepository(r SearchRepository) {
	searchRepository = r
}

func Close() {
	searchRepository.Close()
}

func IndexFeed(ctx context.Context, feed *models.Feed) error {
	return searchRepository.IndexFeed(ctx, feed)
}

func SearchFeeds(ctx context.Context, query string) ([]*models.Feed, error) {
	return searchRepository.SearchFeeds(ctx, query)
}
