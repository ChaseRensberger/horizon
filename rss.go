package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type youtubeFeedResponseXml struct {
	Channel     string `xml:"title"`
	ChannelLink struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Videos []struct {
		Title     string `xml:"title"`
		Published string `xml:"published"`
		Link      struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`

		Group struct {
			Thumbnail struct {
				Url string `xml:"url,attr"`
			} `xml:"http://search.yahoo.com/mrss/ thumbnail"`
		} `xml:"http://search.yahoo.com/mrss/ group"`
	} `xml:"entry"`
}

func getRSS(channelId string) {
	url := fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?channel_id=%s", channelId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed to fetch rss feed")
		return
	}

	var feed youtubeFeedResponseXml
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
