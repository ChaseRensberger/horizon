package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func addTrackedChannel(channelId string, client *mongo.Client, usingFallback bool) (*TrackedChannel, error) {

	trackedChannels, err := getAllTrackedChannels(client)
	if err != nil {
		return nil, err
	}
	for _, trackedChannel := range trackedChannels {
		if trackedChannel.ChannelId == channelId {
			return nil, fmt.Errorf("channel with ID %s is already being tracked", channelId)
		}
	}

	channelSnapshot, err := getCurrentChannelSnapshot(channelId, usingFallback)
	if err != nil {
		return nil, err
	}

	newTrackedChannel := TrackedChannel{
		ChannelId:       channelSnapshot.Items[0].ID,
		ChannelName:     channelSnapshot.Items[0].Snippet.Title,
		ProfileImageURL: channelSnapshot.Items[0].Snippet.Thumbnails["default"].URL,
	}

	collection := client.Database("horizon").Collection("trackedChannels")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, newTrackedChannel)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Inserted new tracked channel with ID: %s", res.InsertedID)

	return &newTrackedChannel, nil
}

func getAllTrackedChannels(client *mongo.Client) ([]TrackedChannel, error) {
	collection := client.Database("horizon").Collection("trackedChannels")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trackedChannels []TrackedChannel
	if err = cursor.All(ctx, &trackedChannels); err != nil {
		return nil, err
	}

	return trackedChannels, nil
}

func getCurrentChannelSnapshot(channelId string, usingFallback bool) (*YoutubeChannelResponse, error) {
	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	requestedParts := strings.Join(usedChannelParts, ",")
	requestUrl := fmt.Sprintf("%s/channels?part=%s&id=%s", youtubeApiUrl, requestedParts, channelId)
	if !usingFallback {
		requestUrl = requestUrl + "&key=" + os.Getenv("YOUTUBE_API_KEY")
	}
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

	responseJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var channelResponse YoutubeChannelResponse
	if err := json.Unmarshal(responseJson, &channelResponse); err != nil {
		return nil, err
	}

	return &channelResponse, nil
}

func getCurrentVideoSnapshot(videoId string, usingFallback bool) (*YoutubeVideoResponse, error) {
	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	requestedParts := strings.Join(usedVideoParts, ",")
	requestUrl := fmt.Sprintf("%s/videos?part=%s&id=%s", youtubeApiUrl, requestedParts, videoId)
	if !usingFallback {
		requestUrl = requestUrl + "&key=" + os.Getenv("YOUTUBE_API_KEY")
	}
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
