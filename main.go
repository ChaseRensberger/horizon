package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("No .env.local file found")
	}
}

func scrapeLukeJ() {
	channelId := "UCYzV77unbAR8KiIoSm4zdUw"
	getRSS(channelId)
}

func main() {
	happenEvery(time.Second*10, scrapeLukeJ)

	// videoId := "3GAAIKDUMLg"

	// err = uploadVideoSnapshot(db, videoId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
