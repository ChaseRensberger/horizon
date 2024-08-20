package main

type Channel struct {
	ChannelID       string `json:"channelId"`
	SubscriberCount int    `json:"subscriberCount"`
	ViewCount       int    `json:"viewCount"`
	VideoCount      int    `json:"videoCount"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type ChannelStatisticsItem struct {
	Kind       string `json:"kind"`
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Statistics struct {
		ViewCount             string `json:"viewCount"`
		SubscriberCount       string `json:"subscriberCount"`
		HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
		VideoCount            string `json:"videoCount"`
	} `json:"statistics"`
}

type ChannelStatisticsResponse struct {
	Kind     string                  `json:"kind"`
	Etag     string                  `json:"etag"`
	PageInfo PageInfo                `json:"pageInfo"`
	Items    []ChannelStatisticsItem `json:"items"`
}
