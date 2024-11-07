package core

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

	"horizon/config"
	"horizon/models"
)

func GetTrackedChannelsByHorizonUserId(horizonUserId string, mongoClient *mongo.Client) ([]models.TrackedChannel, error) {
	collection := mongoClient.Database(config.MongoDatabase).Collection("tracked_channels")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{{Key: "horizonUserId", Value: horizonUserId}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trackedChannels []models.TrackedChannel
	if err = cursor.All(ctx, &trackedChannels); err != nil {
		return nil, err
	}

	return trackedChannels, nil
}

func AddTrackedChannel(channelId string, horizonUserId string, mongoClient *mongo.Client) (*models.TrackedChannel, error) {

	channelSnapshot, err := GetCurrentChannelSnapshot(channelId)
	if err != nil {
		return nil, err
	}

	newTrackedChannel := models.TrackedChannel{
		ChannelId:     channelId,
		ChannelName:   channelSnapshot.Items[0].Snippet.Title,
		HorizonUserId: horizonUserId,
	}

	collection := mongoClient.Database(config.MongoDatabase).Collection("tracked_channels")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, newTrackedChannel)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Inserted new tracked channel for horizon user %s with ID: %s", horizonUserId, res.InsertedID)

	currentChannelSnapshot, err := GetCurrentChannelSnapshot(channelId)
	if err != nil {
		return nil, err
	}

	err = AddChannelSnapshotToDatabase(currentChannelSnapshot, mongoClient)
	if err != nil {
		return nil, err
	}

	return &newTrackedChannel, nil
}

func GetAllTrackedChannels(mongoClient *mongo.Client) ([]models.TrackedChannel, error) {
	collection := mongoClient.Database(config.MongoDatabase).Collection("tracked_channels")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trackedChannels []models.TrackedChannel
	if err = cursor.All(ctx, &trackedChannels); err != nil {
		return nil, err
	}

	return trackedChannels, nil
}

func AddChannelSnapshotToDatabase(channelSnapshot *models.ChannelSnapshot, mongoClient *mongo.Client) error {
	collection := mongoClient.Database(config.MongoDatabase).Collection("channel_snapshots")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, channelSnapshot)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted new channel snapshot with ID: %s", res.InsertedID)

	return nil
}

func GetCurrentChannelSnapshot(channelId string) (*models.ChannelSnapshot, error) {
	youtubeApiUrl := os.Getenv("YOUTUBE_API_URL")
	requestedParts := strings.Join(models.UsedChannelParts, ",")
	requestUrl := fmt.Sprintf("%s/channels?part=%s&id=%s", youtubeApiUrl, requestedParts, channelId)
	if !config.UsingFallback {
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

	var channelSnapshot models.ChannelSnapshot
	if err := json.Unmarshal(responseJson, &channelSnapshot); err != nil {
		return nil, err
	}

	channelSnapshot.RetrievedAt = time.Now()

	return &channelSnapshot, nil
}
