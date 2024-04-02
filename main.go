package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// fetchSubredditPosts retrieves subreddit hot gosts.
func fetchSubredditPosts(subreddit string) (*SubredditResponse, error) {
	s := &http.Client{Timeout: time.Second * 8}
	resp, err := s.Get(fmt.Sprintf("https://www.reddit.com/r/%s/hot.json", subreddit))
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch subreddit posts: %v", err)
	}
	defer resp.Body.Close()

	posts := SubredditResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, fmt.Errorf("Failed to decode JSON response: %v", err)
	}

	return &SubredditResponse{Data: posts.Data}, nil
}

func main() {
	posts, err := fetchSubredditPosts("diablo4")
	if err != nil {
		fmt.Printf("Error fetching subreddit posts: %v\n", err)
		os.Exit(1)
	}

	for _, post := range posts.Data.Children {
		parts := strings.Split(post.Data.Created.String(), ".")
		timestamp, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		t := time.Unix(timestamp, 0)
		fmt.Printf("%s - %v\n", post.Data.Title, t)
	}
}
