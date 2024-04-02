package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	posts, _ := fetchSubredditPosts()

	for _, post := range posts.Data.Children {
		fmt.Println(post.Data.Title)
	}
}

func fetchSubredditPosts() (*SubredditResponse, error) {
	s := &http.Client{Timeout: time.Second * 8}
	resp, err := s.Get("https://www.reddit.com/r/diablo4/hot.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	posts := SubredditResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return &SubredditResponse{Data: posts.Data}, nil
}
