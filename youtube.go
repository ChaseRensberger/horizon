package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func uploadChannelSnapshot(channelid string) error {
	channelSnapshot, err := getCurrentChannelSnapshot(channelid)
	if err != nil {
		return err
	}
	// Open the SQLite database
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO channel_snapshots (channel_id, subscriber_count, view_count, video_count) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the channel snapshot data
	_, err = stmt.Exec(channelid, channelSnapshot.SubscriberCount, channelSnapshot.ViewCount, channelSnapshot.VideoCount)
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

	var channelResponse ChannelResponse
	if err := json.NewDecoder(resp.Body).Decode(&channelResponse); err != nil {
		return nil, err
	}

	channel := &Channel{
		ChannelID:       channelResponse.Items[0].ID,
		SubscriberCount: channelResponse.Items[0].Statistics.SubscriberCount,
		ViewCount:       channelResponse.Items[0].Statistics.ViewCount,
		VideoCount:      channelResponse.Items[0].Statistics.VideoCount,
	}

	return channel, nil

}
