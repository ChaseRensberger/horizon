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

func transformVideoResponseToSnapshot(response *VideoSnapshotResponse) *VideoSnapshot {
	return &VideoSnapshot{
		Kind: response.Kind,
		ETag: response.ETag,
		Item: response.Items[0],
	}
}

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

func getMostRecentVideoSnapshotsByChannelId(channelId string, mongoClient *mongo.Client) ([]VideoSnapshotResponse, error) {
	collection := mongoClient.Database(mongoDatabase).Collection("video_snapshots")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: complete pipeline
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "Items[0].Snippet.channelId", Value: channelId}}}},
		// bson.D{{Key: "$sort", Value: bson.D{{Key: "retrievedAt", Value: -1}}}},
		// bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$items.snippet.resourceId.videoId"}, {Key: "latestSnapshot", Value: bson.D{{Key: "$first", Value: "$$ROOT"}}}}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var videoSnapshots []VideoSnapshotResponse
	if err = cursor.All(ctx, &videoSnapshots); err != nil {
		return nil, err
	}

	return videoSnapshots, nil
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

func addVideoSnapshotToDatabase(videoSnapshot *VideoSnapshotResponse, mongoClient *mongo.Client) error {
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

func getCurrentVideoSnapshot(videoId string) (*VideoSnapshotResponse, error) {
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

	var videoSnapshot VideoSnapshotResponse
	if err := json.Unmarshal(responseJson, &videoSnapshot); err != nil {
		return nil, err
	}

	// videoSnapshot.RetrievedAt = time.Now()
	// videoSnapshot.IsShort = isShort(&videoSnapshot)

	return &videoSnapshot, nil
}

func getCurrentVideoSnapshotAndAddToDatabase(videoId string, mongoClient *mongo.Client) (*VideoSnapshotResponse, error) {
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

func isShort(video *VideoSnapshotResponse) bool {
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
