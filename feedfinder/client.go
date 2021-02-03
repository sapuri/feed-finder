package feedfinder

import (
	"context"
)

type FeedFinder interface {
	FindFeeds(ctx context.Context, siteURL string) (feedURLs []*Feed, err error)
}

type clientImpl struct{}

func (c *clientImpl) FindFeeds(ctx context.Context, siteURL string) ([]*Feed, error) {
	return newfeedFinder().FindFeeds(ctx, siteURL)
}

func New() FeedFinder {
	return &clientImpl{}
}
