package main

import "time"

type PlaylistItemSnapshot struct {
	Kind          string `json:"kind"`
	ETag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	}
	Items []struct {
		Kind           string `json:"kind"`
		ETag           string `json:"etag"`
		ID             string `json:"id"`
		ContentDetails struct {
			VideoID          string `json:"videoId"`
			VideoPublishedAt string `json:"videoPublishedAt"`
		}
	}
	// These are fields that won't come back from the youtube api but are important for our application
	RetrievedAt time.Time `bson:"retrievedAt,omitempty"`
}
