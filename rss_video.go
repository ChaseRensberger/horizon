package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func addRSSSnapshotToDatabase(rssVideoSnapshot *RSSVideoSnapshot, mongoClient *mongo.Client) error {
	collection := mongoClient.Database(mongoDatabase).Collection("rss_video_snapshots")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second
	defer cancel()

	res, err := collection.InsertOne(ctx, rssVideoSnapshot)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted new rss video snapshot with ID: %s", rs.InsertedID)

	return nil
}

func getRecentVideoIdsWithRSS(horizonUserId string, mongoClient *mongo.Client) ([]DerivedRSSVideoSnapshot, error) {
	trackedChannels, err := getTrackedChannelsByHorizonUserId(horizonUserId, mongoClient)
	if err != nil {
		return nil, err
	}

	var derivedRSSVideoSnapshots []DerivedRSSVideoSnapshot
	for _, trackedChannel := range trackedChannels {
		url := fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?channel_id=%s", trackedChannel.ChannelId)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch rss feed: %d", resp.StatusCode)
		}

		var feed RSSVideoSnapshot
		if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
			return nil, err
		}

		videos := make([]struct {
			VideoId   string
			Title     string
			Published string
			Updated   string
			Link      string
			Thumbnail string
			Views     string
			Rating    string
		}, len(feed.Videos))

		for i, v := range feed.Videos {
			videos[i] = struct {
				VideoId   string
				Title     string
				Published string
				Updated   string
				Link      string
				Thumbnail string
				Views     string
				Rating    string
			}{
				VideoId:   v.VideoId,
				Title:     v.Title,
				Published: v.Published,
				Updated:   v.Updated,
				Link:      v.Link.Href,
				Thumbnail: v.Group.Thumbnail.Url,
				Views:     v.Group.Community.Views,
				Rating:    v.Group.Community.StarRating.Average,
			}
		}

		derivedRSSVideoSnapshots = append(derivedRSSVideoSnapshots, DerivedRSSVideoSnapshot{
			ChannelName: feed.ChannelName,
			ChannelId:   trackedChannel.ChannelId,
			ChannelLink: feed.ChannelLink.Href,
			Videos:      videos,
		})
	}

	return derivedRSSVideoSnapshots, nil
}
