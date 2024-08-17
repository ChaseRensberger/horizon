package main

import (
	"fmt"
	"net/http"
	"os"
)

func getCurrentChannelSnapshot(channelId string) (*Channel, error) {

	// fetch current snapshot

	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	part := "statistics"
	requestUrl := fmt.Sprintf("%s/channels?part=%s&channelId=%s", youtubeApiUrl, part, channelId)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil

	// upload to db
}
