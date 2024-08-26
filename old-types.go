package main

// DATABASE TYPES
type Channel struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Country     string `json:"country"`
	CustomURL   string `json:"custom_url"`
}

type Video struct {
	VideoID   string `json:"video_id"`
	ChannelID string `json:"channel_id"`
}

type ChannelSnapshot struct {
	ChannelID       string `json:"channelId"`
	SubscriberCount int    `json:"subscriberCount"`
	ViewCount       int    `json:"viewCount"`
	VideoCount      int    `json:"videoCount"`
}

type VideoSnapshot struct {
	VideoSnapshotID string `json:"video_snapshot_id"`
	VideoID         string `json:"video_id"`
	ViewCount       string `json:"view_count"`
	LikeCount       string `json:"like_count"`
	CommentCount    string `json:"comment_count"`
}

// RESPONSE TYPES

type YoutubeResponse struct {
	Kind     string `json:"kind"`
	ETag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

type YoutubeResponseItems struct {
	Kind string `json:"kind"`
	ETag string `json:"etag"`
	ID   string `json:"id"`
}

type ChannelStatisticsResponse struct {
	YoutubeResponse
	Items []struct {
		YoutubeResponseItems
		Statistics struct {
			ViewCount             string `json:"viewCount"`
			SubscriberCount       string `json:"subscriberCount"`
			HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
			VideoCount            string `json:"videoCount"`
		}
	} `json:"items"`
}

type ChannelSnippetResponse struct {
	YoutubeResponse
	Items []struct {
		YoutubeResponseItems
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			CustomURL   string `json:"customUrl"`
			PublishedAt string `json:"publishedAt"`
			Country     string `json:"country"`
		}
	} `json:"items"`
}

type VideoStatisticsResponse struct {
	YoutubeResponse
	Items []struct {
		YoutubeResponseItems
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		}
	} `json:"items"`
}

type VideoSnippetResponse struct {
	YoutubeResponse
	Items []struct {
		YoutubeResponseItems
		Snippet struct {
			publishedAt          string      `json:"publishedAt"`
			channelID            string      `json:"channelId"`
			title                string      `json:"title"`
			description          string      `json:"description"`
			ChannelTitle         string      `json:"channelTitle"`
			categoryID           string      `json:"categoryId"`
			liveBroadcastContent string      `json:"liveBroadcastContent"`
			defaultAudioLanguage string      `json:"defaultAudioLanguage"`
			thumbnails           []Thumbnail `json:"thumbnails"`
		}
	} `json:"items"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
