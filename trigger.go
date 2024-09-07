package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

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
