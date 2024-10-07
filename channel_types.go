package main

import "time"

// var channelParts = []string{"snippet", "contentDetails", "statistics", "topicDetails", "status", "brandingSettings", "auditDetails", "contentOwnerDetails", "localizations"}
var usedChannelParts = []string{"snippet", "contentDetails", "statistics", "topicDetails", "status", "brandingSettings", "localizations"}

type TrackedChannel struct {
	ChannelId     string `bson:"_id,omitempty"`
	ChannelName   string `bson:"channelName,omitempty"`
	HorizonUserId string `bson:"horizonUserId,omitempty"`
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

type DerivedChannelSnapshot struct {
	Kind string `json:"kind"`
	ETag string `json:"etag"`
	Item struct {
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
