package models

type TrackedChannel struct {
	ChannelId   string `bson:"_id,omitempty"`
	ChannelName string `bson:"channelName,omitempty"`
}
