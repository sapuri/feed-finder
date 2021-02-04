# feed-finder
A Go library for finding RSS feed URLs from the given URL

## Installation
```
go get github.com/sapuri/feed-finder
```

(optional) To run unit tests:
```
make test
```

## Examples
```go
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
	feeds, err := ff.FindFeeds(ctx, siteURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Title)
		fmt.Println(feed.URL)
	}
}
```
