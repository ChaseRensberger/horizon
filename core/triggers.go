package core

// import (
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func uploadTrigger(mongoClient *mongo.Client) error {
// 	trackedChannels, err := GetAllTrackedChannels(mongoClient)
// 	if err != nil {
// 		return err
// 	}

// 	for _, trackedChannel := range trackedChannels {
// 		derivedRSSVideoSnapshots, err := GetRecentVideoIdsWithRSS(trackedChannel.HorizonUserId, mongoClient)
// 		if err != nil {
// 			return err
// 		}
// 		for _, derivedRSSVideoSnapshot := range derivedRSSVideoSnapshots {
// 			for _, video := range derivedRSSVideoSnapshot.Videos {
// 				_, err := addTrackedVideo(video.VideoId, derivedRSSVideoSnapshot.ChannelId, mongoClient)
// 				if err == nil {
// 					// insert won't succeed most of the time
// 					fmt.Printf("Inserted new tracked video with ID: %s", video.VideoId)
// 					_, err = getCurrentVideoSnapshotAndAddToDatabase(video.VideoId, mongoClient)
// 					if err != nil {
// 						return err
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func videoUpdateTrigger(mongoClient *mongo.Client) error {
// 	trackedVideos, err := getAllTrackedVideos(mongoClient)
// 	if err != nil {
// 		return err
// 	}

// 	for _, trackedVideo := range trackedVideos {
// 		videoSnapshot, err := getCurrentVideoSnapshot(trackedVideo.VideoId)
// 		if err != nil {
// 			return err
// 		}

// 		err = addVideoSnapshotToDatabase(videoSnapshot, mongoClient)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func channelUpdateTrigger(mongoClient *mongo.Client) error {
// 	trackedChannels, err := getAllTrackedChannels(mongoClient)
// 	if err != nil {
// 		return err
// 	}

// 	for _, trackedChannel := range trackedChannels {
// 		channelSnapshot, err := getCurrentChannelSnapshot(trackedChannel.ChannelId)
// 		if err != nil {
// 			return err
// 		}

// 		err = addChannelSnapshotToDatabase(channelSnapshot, mongoClient)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
