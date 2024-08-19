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

type ChannelItem struct {
	Kind       string `json:"kind"`
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Statistics struct {
		ViewCount             int  `json:"viewCount"`
		SubscriberCount       int  `json:"subscriberCount"`
		HiddenSubscriberCount bool `json:"hiddenSubscriberCount"`
		VideoCount            int  `json:"videoCount"`
	} `json:"statistics"`
}

type ChannelResponse struct {
	Kind     string        `json:"kind"`
	Etag     string        `json:"etag"`
	PageInfo PageInfo      `json:"pageInfo"`
	Items    []ChannelItem `json:"items"`
}
