package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func uploadChannelSnapshot(db *sql.DB, channelid string) error {
// 	channelSnapshot, err := getCurrentChannelSnapshot(channelid)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	stmt, err := db.Prepare("INSERT INTO channel_snapshots (subscriber_count, view_count, video_count) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(channelSnapshot.SubscriberCount, channelSnapshot.ViewCount, channelSnapshot.VideoCount)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func getCurrentChannelSnapshot(channelId string) (*ChannelSnapshot, error) {

// 	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
// 	part := "statistics"
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

// 	var channelResponse ChannelStatisticsResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&channelResponse); err != nil {
// 		return nil, err
// 	}

// 	channel := &ChannelSnapshot{
// 		ChannelID:       channelResponse.Items[0].ID,
// 		SubscriberCount: convertToInt(channelResponse.Items[0].Statistics.SubscriberCount),
// 		ViewCount:       convertToInt(channelResponse.Items[0].Statistics.ViewCount),
// 		VideoCount:      convertToInt(channelResponse.Items[0].Statistics.VideoCount),
// 	}

// 	return channel, nil

// }

// func createChannelFromId(db *sql.DB, channelId string) error {
// 	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
// 	part := "snippet"
// 	requestUrl := fmt.Sprintf("%s/channels?part=%s&id=%s", youtubeApiUrl, part, channelId)
// 	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Accept", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to fetch channel data: %d", resp.StatusCode)
// 	}

// 	var channelSnippetResponse ChannelSnippetResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&channelSnippetResponse); err != nil {
// 		return err
// 	}

// 	channel := &Channel{
// 		ChannelID:   channelSnippetResponse.Items[0].ID,
// 		ChannelName: channelSnippetResponse.Items[0].Snippet.Title,
// 		Country:     channelSnippetResponse.Items[0].Snippet.Country,
// 		CustomURL:   channelSnippetResponse.Items[0].Snippet.CustomURL,
// 	}

// 	stmt, err := db.Prepare("INSERT INTO channels (channel_id, channel_name, country, custom_url) VALUES (?, ?, ?, ?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(channel.ChannelID, channel.ChannelName, channel.Country, channel.CustomURL)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

// func uploadVideoSnapshot(db *sql.DB, videoId string) error {
// 	videoSnapshot, err := getCurrentVideoSnapshot(videoId)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	stmt, err := db.Prepare("INSERT INTO video_snapshots (view_count, like_count, comment_count) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(videoSnapshot.ViewCount, videoSnapshot.LikeCount, videoSnapshot.CommentCount)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func getCurrentVideoSnapshot(videoId string) (*YoutubeVideoResponse, error) {
	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	allParts := strings.Join(usedVideoParts, ",")
	requestUrl := fmt.Sprintf("%s/videos?part=%s&id=%s", youtubeApiUrl, allParts, videoId)
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
		return nil, fmt.Errorf("failed to fetch video data: %d", resp.StatusCode)
	}

	responseJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var videoResponse YoutubeVideoResponse
	if err := json.Unmarshal(responseJson, &videoResponse); err != nil {
		return nil, err
	}

	return &videoResponse, nil
}
