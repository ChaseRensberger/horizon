package main

type Channel struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Country     string `json:"country"`
	CustomURL   string `json:"custom_url"`
}

type ChannelSnapshot struct {
	ChannelID       string `json:"channelId"`
	SubscriberCount int    `json:"subscriberCount"`
	ViewCount       int    `json:"viewCount"`
	VideoCount      int    `json:"videoCount"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
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
