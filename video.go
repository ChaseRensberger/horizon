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

func addTrackedVideo(videoId string, channelId string, mongoClient *mongo.Client) (*TrackedVideo, error) {
	// may need to insert some other video metadata
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

func getMostRecentVideoSnapshotsByChannelId(channelId string, mongoClient *mongo.Client) ([]VideoSnapshot, error) {
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

	var videoSnapshots []VideoSnapshot
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

	// videoSnapshot.RetrievedAt = time.Now()
	// videoSnapshot.IsShort = isShort(&videoSnapshot)

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
