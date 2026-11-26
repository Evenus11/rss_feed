package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/evenus11/rss_feed/internal/rss"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: feedwatch <fetch|new>")
		os.Exit(2)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	feed := rss.Feed{}

	args := os.Args[2:]

	var err error

	switch os.Args[1] {
	case "fetch":
		feed, err = rss.Fetch(client, args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching feed: %v\n", err)
			os.Exit(1)
		}
	case "new":
		err = newFeed(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating new feed: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Usage: feedwatch <fetch|new>")
		os.Exit(2)
	}

	fmt.Println(feed.Title)
	for _, item := range feed.Items {
		fmt.Printf("- %s\n  %s\n", item.Title, item.Link)
	}
}
