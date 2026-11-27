package rss

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Title   string   `xml:"channel>title"`
	Items   []Item   `xml:"channel>item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	GUID  string `xml:"guid"`
}

func Fetch(Client *http.Client, url string) (Feed, error) {

	var feed Feed

	resp, err := Client.Get(url)
	if err != nil {
		return feed, fmt.Errorf("failed to fetch feed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return feed, fmt.Errorf("failed to fetch feed: status code %d", resp.StatusCode)
	}

	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return feed, fmt.Errorf("failed to decode feed: %v", err)
	}

	return feed, nil

}

func NewFeed(url string) error {
	// Placeholder for creating a new feed
	// In a real application, you might want to store the feed URL in a database or a file
	if url == "" {
		return fmt.Errorf("feed URL cannot be empty")
	}
	fmt.Printf("New feed created: %s\n", url)
	return nil
}
