package feedfinder

import (
	"context"
)

type FeedFinder interface {
	FindFeeds(ctx context.Context, siteURL string) ([]*Feed, error)
}

type clientImpl struct{}

func (c *clientImpl) FindFeeds(ctx context.Context, siteURL string) ([]*Feed, error) {
	return newFeedFinder().FindFeeds(ctx, siteURL)
}

func New() FeedFinder {
	return &clientImpl{}
}
