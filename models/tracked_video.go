package models

type TrackedVideo struct {
	VideoId   string `bson:"_id,omitempty"`
	ChannelId string `bson:"channelId,omitempty"`
}
