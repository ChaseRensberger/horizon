package models

type RSSVideoSnapshot struct {
	ChannelName string `xml:"title"`
	ChannelId   string `xml:"yt:channelId"`
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
