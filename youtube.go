package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getCurrentChannelSnapshot(channelId string) (*Channel, error) {

	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	part := "statistics"
	requestUrl := fmt.Sprintf("%s/channels?part=%s&id=%s", youtubeApiUrl, part, channelId)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch channel data: %d", resp.StatusCode)
	}

	var channel Channel
	if err := json.NewDecoder(resp.Body).Decode(&channel); err != nil {
		return nil, err
	}

	return &channel, nil

}
