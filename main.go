package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("No .env.local file found")
	}
}

func main() {
	channelId := "UCYzV77unbAR8KiIoSm4zdUw"
	err := uploadChannelSnapshot(channelId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Channel snapshot uploaded successfully")
}
