package main

import (
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	subreddit := "diablo4"
	// Create a wait group to wait for both goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Channels to communicate results and errors
	postsCh := make(chan *HotPostsResponse)
	aboutCh := make(chan *AboutResponse)
	errCh := make(chan error)

	// Fetch hot posts concurrently
	go func() {
		defer wg.Done()
		posts, err := FetchHotPosts(subreddit)
		if err != nil {
			errCh <- fmt.Errorf("Error fetching subreddit posts: %v", err)
			return
		}
		postsCh <- posts
	}()

	// Fetch subreddit about concurrently
	go func() {
		defer wg.Done()
		about, err := FetchAbout(subreddit)
		if err != nil {
			errCh <- fmt.Errorf("Error fetching subreddit about: %v", err)
			return
		}
		aboutCh <- about
	}()

	// Wait for both goroutines to finish
	go func() {
		wg.Wait()
		close(postsCh)
		close(aboutCh)
		close(errCh)
	}()

	// Handle results and errors
	var posts *HotPostsResponse
	var about *AboutResponse
	var err error
	for {
		select {
		case posts = <-postsCh:
			// Process posts
			for _, post := range posts.Data.Children {
				parts := strings.Split(post.Data.Created.String(), ".")
				timestamp, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					fmt.Printf("Error parsing timestamp: %v\n", err)
					continue
				}
				t := time.Unix(timestamp, 0)
				fmt.Printf("%s - %v\n", post.Data.Title, t)
			}
		case about = <-aboutCh:
			// Process about
			fmt.Printf("\n%s\n%s\n", about.Data.Title, about.Data.PublicDescription)
		case err = <-errCh:
			fmt.Println(err)
			os.Exit(1)
		}
		// Exit loop when both posts and about are processed
		if posts != nil && about != nil {
			break
		}
	}
}
