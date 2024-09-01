package main

import (
	"time"
)

// var channelParts = []string{"snippet", "contentDetails", "statistics", "topicDetails", "status", "brandingSettings", "auditDetails", "contentOwnerDetails", "localizations"}
// var videoParts = []string{"snippet", "contentDetails", "status", "statistics", "player", "topicDetails", "recordingDetails", "fileDetails", "processingDetails", "suggestions", "liveStreamingDetails", "localizations"}
var usedChannelParts = []string{"snippet", "contentDetails", "statistics", "topicDetails", "status", "brandingSettings", "localizations"}
var usedVideoParts = []string{"snippet", "contentDetails", "status", "statistics", "player", "topicDetails", "liveStreamingDetails", "localizations"}

type TrackedChannel struct {
	ChannelId string `bson:"_id,omitempty"`
}

type TrackedVideo struct {
	VideoId string `bson:"_id,omitempty"`
	// ChannelId string `bson:"channelId,omitempty"`
}

type ChannelSnapshot struct {
	Kind  string `json:"kind"`
	ETag  string `json:"etag"`
	Items []struct {
		Kind    string `json:"kind"`
		ETag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			CustomURL   string `json:"customUrl"`
			PublishedAt string `json:"publishedAt"`
			Thumbnails  map[string]struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			}
			DefaultLanguage string `json:"defaultLanguage"`
			Localized       struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}
			Country string `json:"country"`
		}
		ContentDetails struct {
			RelatedPlaylists struct {
				Likes     string `json:"likes"`
				Favorites string `json:"favorites"`
				Uploads   string `json:"uploads"`
			}
		}
		Statistics struct {
			ViewCount             string `json:"viewCount"`
			SubscriberCount       string `json:"subscriberCount"`
			HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
			VideoCount            string `json:"videoCount"`
		}
		TopicDetails struct {
			TopicIds        []string `json:"topicIds"`
			TopicCategories []string `json:"topicCategories"`
		}
		Status struct {
			PrivacyStatus           string `json:"privacyStatus"`
			IsLinked                bool   `json:"isLinked"`
			LongUploadsStatus       string `json:"longUploadsStatus"`
			MadeForKids             bool   `json:"madeForKids"`
			SelfDeclaredMadeForKids bool   `json:"selfDeclaredMadeForKids"`
		}
		BrandingSettings struct {
			Channel struct {
				Title                    string `json:"title"`
				Description              string `json:"description"`
				Keywords                 string `json:"keywords"`
				TrackingAnalyticsAccount string `json:"trackingAnalyticsAccount"`
				UnsubscribedTrailer      string `json:"unsubscribedTrailer"`
				DefaultLanguage          string `json:"defaultLanguage"`
				Country                  string `json:"country"`
			}
			Watch struct {
				TextColor          string `json:"textColor"`
				BackgroundColor    string `json:"backgroundColor"`
				FeaturedPlaylistID string `json:"featuredPlaylistId"`
			}
		}
		// AuditDetails struct {
		// 	OverallGoodStanding             bool `json:"overallGoodStanding"`
		// 	CommunityGuidelinesGoodStanding bool `json:"communityGuidelinesGoodStanding"`
		// 	CopyrightGoodStanding           bool `json:"copyrightGoodStanding"`
		// 	ContentIdClaimsGoodStanding     bool `json:"contentIdClaimsGoodStanding"`
		// }
		// ContentOwnerDetails struct {
		// 	ContentOwner string `json:"contentOwner"`
		// 	TimeLinked   string `json:"timeLinked"`
		// }
		Localizations map[string]struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
	}
	// These are fields that won't come back from the youtube api but are important for our application
	RetrievedAt time.Time `bson:"retrievedAt,omitempty"`
}

type VideoSnapshot struct {
	Kind  string `json:"kind"`
	ETag  string `json:"etag"`
	Items []struct {
		Kind    string `json:"kind"`
		ETag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt string `json:"publishedAt"`
			ChannelID   string `json:"channelId"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Thumbnails  map[string]struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			}
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			DefaultLanguage      string   `json:"defaultLanguage"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}
			DefaultAudioLanguage string `json:"defaultAudioLanguage"`
		}
		ContentDetails struct {
			Duration          string `json:"duration"`
			Dimension         string `json:"dimension"`
			Definition        string `json:"definition"`
			Caption           string `json:"caption"`
			LicensedContent   bool   `json:"licensedContent"`
			RegionRestriction struct {
				Allowed []string `json:"allowed"`
				Blocked []string `json:"blocked"`
			}
			ContentRating      map[string]string `json:"contentRating"`
			Projection         string            `json:"projection"`
			HasCustomThumbnail bool              `json:"hasCustomThumbnail"`
		}
		Status struct {
			UploadStatus            string `json:"uploadStatus"`
			FailureReason           string `json:"failureReason"`
			RejectionReason         string `json:"rejectionReason"`
			PrivacyStatus           string `json:"privacyStatus"`
			PublishAt               string `json:"publishAt"`
			License                 string `json:"license"`
			Embeddable              bool   `json:"embeddable"`
			PublicStatsViewable     bool   `json:"publicStatsViewable"`
			MadeForKids             bool   `json:"madeForKids"`
			SelfDeclaredMadeForKids bool   `json:"selfDeclaredMadeForKids"`
		}
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			DislikeCount  string `json:"dislikeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		}
		Player struct {
			EmbedHTML   string `json:"embedHtml"`
			EmbedHeight int    `json:"embedHeight"`
			EmbedWidth  int    `json:"embedWidth"`
		}
		TopicDetails struct {
			TopicIds         []string `json:"topicIds"`
			RelevantTopicIds []string `json:"relevantTopicIds"`
			TopicCategories  []string `json:"topicCategories"`
		}
		// RecordingDetails struct {
		// 	RecordingDate string `json:"recordingDate"`
		// }
		// FileDetails struct {
		// 	FileName     string `json:"fileName"`
		// 	FileSize     int    `json:"fileSize"`
		// 	FileType     string `json:"fileType"`
		// 	Container    string `json:"container"`
		// 	VideoStreams []struct {
		// 		WidthPixels  int     `json:"widthPixels"`
		// 		HeightPixels int     `json:"heightPixels"`
		// 		FrameRateFps float64 `json:"frameRateFps"`
		// 		AspectRatio  float64 `json:"aspectRatio"`
		// 		Codec        string  `json:"codec"`
		// 		BitrateBps   string  `json:"bitrateBps"`
		// 		Rotation     string  `json:"rotation"`
		// 		Vendor       string  `json:"vendor"`
		// 	}
		// 	AudioStreams []struct {
		// 		ChannelCount int    `json:"channelCount"`
		// 		Codec        string `json:"codec"`
		// 		BitrateBps   string `json:"bitrateBps"`
		// 		Vendor       string `json:"vendor"`
		// 	}
		// 	DurationMs   string `json:"durationMs"`
		// 	BitrateBps   string `json:"bitrateBps"`
		// 	CreationTime string `json:"creationTime"`
		// }
		// ProcessingDetails struct {
		// 	ProcessingStatus   string `json:"processingStatus"`
		// 	ProcessingProgress struct {
		// 		PartsTotal     int    `json:"partsTotal"`
		// 		PartsProcessed int    `json:"partsProcessed"`
		// 		TimeLeftMs     string `json:"timeLeftMs"`
		// 	}
		// 	ProcessingFailureReason       string `json:"processingFailureReason"`
		// 	FileDetailsAvailability       string `json:"fileDetailsAvailability"`
		// 	ProcessingIssuesAvailability  string `json:"processingIssuesAvailability"`
		// 	TagSuggestionsAvailability    string `json:"tagSuggestionsAvailability"`
		// 	EditorSuggestionsAvailability string `json:"editorSuggestionsAvailability"`
		// 	ThumbnailsAvailability        string `json:"thumbnailsAvailability"`
		// }
		// Suggestions struct {
		// 	ProcessingErrors   []string `json:"processingErrors"`
		// 	ProcessingWarnings []string `json:"processingWarnings"`
		// 	ProcessingHints    []string `json:"processingHints"`
		// 	TagSuggestions     []struct {
		// 		Tag               string   `json:"tag"`
		// 		CategoryRestricts []string `json:"categoryRestricts"`
		// 	}
		// 	EditorSuggestions []string `json:"editorSuggestions"`
		// }
		LiveStreamingDetails struct {
			ActualStartTime    string `json:"actualStartTime"`
			ActualEndTime      string `json:"actualEndTime"`
			ScheduledStartTime string `json:"scheduledStartTime"`
			ScheduledEndTime   string `json:"scheduledEndTime"`
			ConcurrentViewers  string `json:"concurrentViewers"`
			ActiveLiveChatID   string `json:"activeLiveChatId"`
		}
		Localizations map[string]struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
	}
	// These are fields that won't come back from the youtube api but are important for our application
	RetrievedAt time.Time `bson:"retrievedAt,omitempty"`
	IsShort     bool      `bson:"isShort,omitempty"`
}

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

type RSSVideoSnapshot struct {
	Channel     string `xml:"title"`
	ChannelLink struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Videos []struct {
		VideoId   string `xml:"videoId"`
		Title     string `xml:"title"`
		Published string `xml:"published"`
		Updated   string `xml:"updated"`
		Link      struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`

		Group struct {
			Thumbnail struct {
				Url string `xml:"url,attr"`
			} `xml:"http://search.yahoo.com/mrss/ thumbnail"`
			Community struct {
				StarRating struct {
					Average string `xml:"average,attr"`
					Count   string `xml:"count,attr"`
				} `xml:"http://search.yahoo.com/mrss/ starRating"`
				Views string `xml:"http://search.yahoo.com/mrss/ views"`
			} `xml:"http://search.yahoo.com/mrss/ community"`
		} `xml:"http://search.yahoo.com/mrss/ group"`
	} `xml:"entry"`
}
