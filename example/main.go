package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sapuri/feed-finder/feedfinder"
)

func main() {
	ctx := context.Background()
	const siteURL = "https://www3.nhk.or.jp/news/"

	ff := feedfinder.New()
	feedURLs, err := ff.FindFeeds(ctx, siteURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(feedURLs)
}
