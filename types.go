package main

import "encoding/json"

// HOT POSTS
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

type HotPostsResponse struct {
	Data struct {
		Children []Children `json:"children"`
	} `json:"data"`
}

// ABOUT
type AboutResponse struct {
	Data struct {
		Title                 string      `json:"title"`
		PrimaryColor          string      `json:"primary_color"`
		BannerBackgroundColor string      `json:"banner_background_color"`
		IconIMG               string      `json:"icon_img"`
		DisplayNamePrefixed   string      `json:"display_name_prefixed"`
		AccountsActive        json.Number `json:"accounts_active"`
		PublicDescription     string      `json:"public_description"`
	} `json:"data"`
}
