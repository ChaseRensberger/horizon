package main

// DATABASE TYPES
type Channel struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Country     string `json:"country"`
	CustomURL   string `json:"custom_url"`
}

type Video struct {
	VideoID   string `json:"video_id"`
	ChannelID string `json:"channel_id"`
}

type ChannelSnapshot struct {
	ChannelID       string `json:"channelId"`
	SubscriberCount int    `json:"subscriberCount"`
	ViewCount       int    `json:"viewCount"`
	VideoCount      int    `json:"videoCount"`
}

type VideoSnapshot struct {
	VideoSnapshotID string `json:"video_snapshot_id"`
	VideoID         string `json:"video_id"`
	ViewCount       string `json:"view_count"`
	LikeCount       string `json:"like_count"`
	CommentCount    string `json:"comment_count"`
}

// RESPONSE TYPES

type ChannelStatisticsResponse struct {
	Kind     string                 `json:"kind"`
	Etag     string                 `json:"etag"`
	PageInfo PageInfo               `json:"pageInfo"`
	Items    []ChannelStatisticItem `json:"items"`
}

type ChannelSnippetResponse struct {
	Kind     string               `json:"kind"`
	Etag     string               `json:"etag"`
	PageInfo PageInfo             `json:"pageInfo"`
	Items    []ChannelSnippetItem `json:"items"`
}

type VideoStatisticsResponse struct {
	Kind     string               `json:"kind"`
	Etag     string               `json:"etag"`
	PageInfo PageInfo             `json:"pageInfo"`
	Items    []VideoStatisticItem `json:"items"`
}

type VideoSnippetResponse struct {
	Kind     string             `json:"kind"`
	Etag     string             `json:"etag"`
	PageInfo PageInfo           `json:"pageInfo"`
	Items    []VideoSnippetItem `json:"items"`
}

type ChannelStatisticItem struct {
	Kind       string           `json:"kind"`
	Etag       string           `json:"etag"`
	ID         string           `json:"id"`
	Statistics ChannelStatistic `json:"statistics"`
}

type ChannelSnippetItem struct {
	Kind    string         `json:"kind"`
	Etag    string         `json:"etag"`
	ID      string         `json:"id"`
	Snippet ChannelSnippet `json:"snippet"`
}

type VideoStatisticItem struct {
	Kind       string         `json:"kind"`
	Etag       string         `json:"etag"`
	ID         string         `json:"id"`
	Statistics VideoStatistic `json:"statistics"`
}

type VideoSnippetItem struct {
	Kind    string       `json:"kind"`
	Etag    string       `json:"etag"`
	ID      string       `json:"id"`
	Snippet VideoSnippet `json:"snippet"`
}

type ChannelStatistic struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

type ChannelSnippet struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CustomURL   string `json:"customUrl"`
	PublishedAt string `json:"publishedAt"`
	Country     string `json:"country"`
}

type VideoStatistic struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

type VideoSnippet struct {
	publishedAt          string `json:"publishedAt"`
	channelID            string `json:"channelId"`
	title                string `json:"title"`
	description          string `json:"description"`
	ChannelTitle         string `json:"channelTitle"`
	categoryID           string `json:"categoryId"`
	liveBroadcastContent string `json:"liveBroadcastContent"`
	defaultAudioLanguage string `json:"defaultAudioLanguage"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}
