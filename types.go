package main

type Channel struct {
	ChannelID  string `json:"channelId"`
	Statistics struct {
		SubscriberCount string `json:"subscriberCount"`
		ViewCount       string `json:"viewCount"`
		VideoCount      string `json:"videoCount"`
	} `json:"statistics"`
}
