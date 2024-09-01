package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func addTrackedVideo(videoId string, channelId string, mongoClient *mongo.Client) (*TrackedVideo, error) {
	// trackedVideos, err := getAllTrackedVideos(client)
	// if err != nil {
	// 	return nil, err
	// }
	// for _, trackedVideo := range trackedVideos {
	// 	if trackedVideo.VideoId == videoId {
	// 		return nil, fmt.Errorf("video with ID %s is already being tracked", videoId)
	// 	}
	// }
	//
	// videoSnapshot, err := getCurrentVideoSnapshot(videoId)
	// if err != nil {
	// 	return nil, err
	// }

	newTrackedVideo := TrackedVideo{
		VideoId:   videoId,
		ChannelId: channelId,
	}

	collection := mongoClient.Database(mongoDatabase).Collection("tracked_videos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, newTrackedVideo)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Inserted new tracked video with ID: %s", res.InsertedID)

	return &newTrackedVideo, nil
}

func addTrackedChannel(channelId string, mongoClient *mongo.Client) (*TrackedChannel, error) {

	channelSnapshot, err := getCurrentChannelSnapshot(channelId)
	if err != nil {
		return nil, err
	}

	newTrackedChannel := TrackedChannel{
		ChannelId:   channelId,
		ChannelName: channelSnapshot.Items[0].Snippet.Title,
	}

	collection := mongoClient.Database(mongoDatabase).Collection("tracked_channels")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, newTrackedChannel)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Inserted new tracked channel with ID: %s", res.InsertedID)

	currentChannelSnapshot, err := getCurrentChannelSnapshot(channelId)
	if err != nil {
		return nil, err
	}

	err = addChannelSnapshotToDatabase(currentChannelSnapshot, mongoClient)
	if err != nil {
		return nil, err
	}

	return &newTrackedChannel, nil
}

func getAllTrackedChannels(mongoClient *mongo.Client) ([]TrackedChannel, error) {
	collection := mongoClient.Database(mongoDatabase).Collection("tracked_channels")
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

func getAllTrackedVideos(mongoClient *mongo.Client) ([]TrackedVideo, error) {
	collection := mongoClient.Database(mongoDatabase).Collection("tracked_videos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trackedVideos []TrackedVideo
	if err = cursor.All(ctx, &trackedVideos); err != nil {
		return nil, err
	}

	return trackedVideos, nil
}

func addChannelSnapshotToDatabase(channelSnapshot *ChannelSnapshot, mongoClient *mongo.Client) error {
	collection := mongoClient.Database(mongoDatabase).Collection("channel_snapshots")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, channelSnapshot)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted new channel snapshot with ID: %s", res.InsertedID)

	return nil
}

func addVideoSnapshotToDatabase(videoSnapshot *VideoSnapshot, mongoClient *mongo.Client) error {
	collection := mongoClient.Database(mongoDatabase).Collection("video_snapshots")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, videoSnapshot)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted new video snapshot with ID: %s", res.InsertedID)

	return nil
}

func getCurrentChannelSnapshot(channelId string) (*ChannelSnapshot, error) {
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
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
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

	var channelSnapshot ChannelSnapshot
	if err := json.Unmarshal(responseJson, &channelSnapshot); err != nil {
		return nil, err
	}

	channelSnapshot.RetrievedAt = time.Now()

	return &channelSnapshot, nil
}

func getCurrentVideoSnapshot(videoId string) (*VideoSnapshot, error) {
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
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
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

	var videoSnapshot VideoSnapshot
	if err := json.Unmarshal(responseJson, &videoSnapshot); err != nil {
		return nil, err
	}

	videoSnapshot.RetrievedAt = time.Now()
	videoSnapshot.IsShort = isShort(&videoSnapshot)

	return &videoSnapshot, nil
}

func getCurrentVideoSnapshotAndAddToDatabase(videoId string, mongoClient *mongo.Client) (*VideoSnapshot, error) {
	videoSnapshot, err := getCurrentVideoSnapshot(videoId)
	if err != nil {
		return nil, err
	}

	err = addVideoSnapshotToDatabase(videoSnapshot, mongoClient)
	if err != nil {
		return nil, err
	}

	return videoSnapshot, nil
}

// unused for the time being but a function like this may be useful in the future
func getRecentVideoIdsFromChannel(channelId string, numVideos int) ([]string, error) {
	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	// this may be a heuristic, but I assume it is always correct
	playlistId := "UU" + channelId[2:]
	requestUrl := fmt.Sprintf("%s/playlistItems?part=contentDetails&playlistId=%s&maxResults=%d&order=date", youtubeApiUrl, playlistId, numVideos)

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch playlist data: %d", resp.StatusCode)
	}

	responseJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var playlistItemSnapshot PlaylistItemSnapshot
	if err := json.Unmarshal(responseJson, &playlistItemSnapshot); err != nil {
		return nil, err
	}

	var videoIds []string

	for _, item := range playlistItemSnapshot.Items {
		videoIds = append(videoIds, item.ContentDetails.VideoID)
	}

	return videoIds, nil
}

func isShort(video *VideoSnapshot) bool {
	return strings.Contains(video.Items[0].ContentDetails.Duration, "M")
}

func getRecentVideoIdsWithRSS(channelId string) ([]VideoIdWithChannel, error) {
	url := fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?channel_id=%s", channelId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed to fetch rss feed")
		return nil, fmt.Errorf("failed to fetch rss feed: %d", resp.StatusCode)
	}

	var feed RSSVideoSnapshot
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		fmt.Println(err)
		return nil, err
	}

	var recentVideoIdsWithChannel []VideoIdWithChannel
	channelName := feed.ChannelName
	for _, video := range feed.Videos {
		recentVideoIdsWithChannel = append(recentVideoIdsWithChannel, VideoIdWithChannel{VideoId: video.VideoId, ChannelId: channelId, ChannelName: channelName})
	}

	return recentVideoIdsWithChannel, nil
}

func uploadTrigger(mongoClient *mongo.Client) error {
	trackedChannels, err := getAllTrackedChannels(mongoClient)
	if err != nil {
		return err
	}

	for _, trackedChannel := range trackedChannels {
		recentVideoIdsWithChannel, err := getRecentVideoIdsWithRSS(trackedChannel.ChannelId)
		if err != nil {
			return err
		}
		for _, videoIdWithChannel := range recentVideoIdsWithChannel {
			// insert won't succeed most of the time
			_, err := addTrackedVideo(videoIdWithChannel.VideoId, videoIdWithChannel.ChannelId, mongoClient)
			if err == nil {
				fmt.Printf("Inserted new tracked video with ID: %s", videoIdWithChannel.VideoId)
				_, err = getCurrentVideoSnapshotAndAddToDatabase(videoIdWithChannel.VideoId, mongoClient)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func videoUpdateTrigger(mongoClient *mongo.Client) error {
	trackedVideos, err := getAllTrackedVideos(mongoClient)
	if err != nil {
		return err
	}

	for _, trackedVideo := range trackedVideos {
		videoSnapshot, err := getCurrentVideoSnapshot(trackedVideo.VideoId)
		if err != nil {
			return err
		}

		err = addVideoSnapshotToDatabase(videoSnapshot, mongoClient)
		if err != nil {
			return err
		}
	}

	return nil
}

func channelUpdateTrigger(mongoClient *mongo.Client) error {
	trackedChannels, err := getAllTrackedChannels(mongoClient)
	if err != nil {
		return err
	}

	for _, trackedChannel := range trackedChannels {
		channelSnapshot, err := getCurrentChannelSnapshot(trackedChannel.ChannelId)
		if err != nil {
			return err
		}

		err = addChannelSnapshotToDatabase(channelSnapshot, mongoClient)
		if err != nil {
			return err
		}
	}

	return nil
}
