package feedfinder

import (
	"context"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/sapuri/feed-finder/feedfinder/errors"
)

type Feed struct {
	Title string
	URL   *url.URL
}

type feedFinder struct{}

func newFeedFinder() *feedFinder {
	return &feedFinder{}
}

func (ff *feedFinder) FindFeeds(ctx context.Context, siteURL string) ([]*Feed, error) {
	// TODO: allow the HTTP option to be specified
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, siteURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		err = errors.NewHTTPError(res.StatusCode)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	links := doc.Find(`link[type="application/rss+xml"]`)
	feeds := make([]*Feed, 0, len(links.Nodes))
	links.Each(func(i int, s *goquery.Selection) {
		var feed Feed
		var err error
		if rawurl, exists := s.Attr("href"); exists {
			if feed.URL, err = url.Parse(rawurl); err != nil {
				return
			}
		}

		if title, exists := s.Attr("title"); exists {
			feed.Title = title
		}

		feeds[i] = &feed
	})

	// TODO: Add more find logic

	return feeds, nil
}
