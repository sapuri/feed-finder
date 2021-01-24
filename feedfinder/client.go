package feedfinder

import (
	"context"

	"github.com/sapuri/feed-finder/feedfinder/internal/feedfinder"
)

type FeedFinder interface {
	FindFeeds(ctx context.Context, siteURL string) (feedURLs []string, err error)
}

type clientImpl struct{}

func (c *clientImpl) FindFeeds(ctx context.Context, siteURL string) (feedURLs []string, err error) {
	return feedfinder.NewFeedFinder().FindFeeds(ctx, siteURL)
}

func New() FeedFinder {
	return &clientImpl{}
}
