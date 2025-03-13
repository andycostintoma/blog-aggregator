package internal

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoded := &RSSFeed{}
	err = xml.NewDecoder(res.Body).Decode(decoded)
	if err != nil {
		return nil, err
	}

	decoded.Channel.Title = html.UnescapeString(decoded.Channel.Title)
	decoded.Channel.Description = html.UnescapeString(decoded.Channel.Description)
	for i, item := range decoded.Channel.Item {
		decoded.Channel.Item[i].Title = html.UnescapeString(item.Title)
		decoded.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return decoded, nil
}
