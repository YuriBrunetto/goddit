package main

import "encoding/json"

type Children struct {
	Data ChildData `json:"data"`
}

type ChildData struct {
	Title                 string      `json:"title"`
	SubredditNamePrefixed string      `json:"subreddit_name_prefixed"`
	Name                  string      `json:"name"`
	LinkFlairText         string      `json:"link_flair_text"`
	LinkFlairBgColor      string      `json:"link_flair_background_color"`
	Thumbnail             string      `json:"thumbnail"`
	Permalink             string      `json:"permalink"`
	URL                   string      `json:"url"`
	Selftext              string      `json:"selftext,omitempty"`
	Author                string      `json:"author"`
	Created               json.Number `json:"created"`
}

type SubredditResponse struct {
	Data struct {
		Children []Children `json:"children"`
	} `json:"data"`
}
