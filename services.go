package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func fetchData(subreddit, endpoint string, response interface{}) error {
	s := &http.Client{Timeout: time.Second * 8}
	resp, err := s.Get(fmt.Sprintf("https://www.reddit.com/r/%s/%s.json", subreddit, endpoint))
	if err != nil {
		return fmt.Errorf("Failed to fetch: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("Failed to decode JSON response: %v", err)
	}

	return nil
}

// FetchHotPosts returns the hot posts of a subreddit
func FetchHotPosts(subreddit string) (*HotPostsResponse, error) {
	var posts HotPostsResponse
	err := fetchData(subreddit, "hot", &posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

// FetchAbout returns the about content of a subreddit
func FetchAbout(subreddit string) (*AboutResponse, error) {
	var about AboutResponse
	err := fetchData(subreddit, "about", &about)
	if err != nil {
		return nil, err
	}
	return &about, nil
}
