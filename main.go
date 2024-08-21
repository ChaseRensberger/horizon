package main

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("No .env.local file found")
	}
}

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	channelId := "UCYzV77unbAR8KiIoSm4zdUw"

	createChannelFromId(db, channelId)

}
