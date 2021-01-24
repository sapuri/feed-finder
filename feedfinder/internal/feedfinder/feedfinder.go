package feedfinder

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/sapuri/feed-finder/feedfinder/errors"
)

type FeedFinder struct{}

func NewFeedFinder() *FeedFinder {
	return &FeedFinder{}
}

func (ff *FeedFinder) FindFeeds(ctx context.Context, siteURL string) (feedURLs []string, err error) {
	// TODO: allow the HTTP option to be specified
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, siteURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		err = errors.NewHTTPError(res.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	links := doc.Find(`link[type="application/rss+xml"]`)
	links.Each(func(i int, s *goquery.Selection) {
		feedURL, exists := s.Attr("href")
		if exists {
			feedURLs = append(feedURLs, feedURL)
		}
	})

	// TODO: Add more find logic

	return
}
