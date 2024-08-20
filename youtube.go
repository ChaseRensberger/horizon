package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func uploadChannelSnapshot(channelid string) error {
	channelSnapshot, err := getCurrentChannelSnapshot(channelid)
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO channel_snapshots (subscriber_count, view_count, video_count) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(channelSnapshot.SubscriberCount, channelSnapshot.ViewCount, channelSnapshot.VideoCount)
	if err != nil {
		return err
	}

	return nil
}

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

	var channelResponse ChannelStatisticsResponse
	if err := json.NewDecoder(resp.Body).Decode(&channelResponse); err != nil {
		return nil, err
	}

	channel := &Channel{
		ChannelID:       channelResponse.Items[0].ID,
		SubscriberCount: convertToInt(channelResponse.Items[0].Statistics.SubscriberCount),
		ViewCount:       convertToInt(channelResponse.Items[0].Statistics.ViewCount),
		VideoCount:      convertToInt(channelResponse.Items[0].Statistics.VideoCount),
	}

	return channel, nil

}

// func createChannelFromId(channelId string) (*Channel, error) {
// 	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
// 	part := "snippet"
// 	requestUrl := fmt.Sprintf("%s/channels?part=%s&id=%s", youtubeApiUrl, part, channelId)
// 	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Accept", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("failed to fetch channel data: %d", resp.StatusCode)
// 	}

// }

func convertToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}
