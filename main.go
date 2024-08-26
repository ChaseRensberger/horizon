package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("No .env.local file found")
	}
}

// func scrapeLukeJ() {
// 	channelId := "UCYzV77unbAR8KiIoSm4zdUw"
// 	getRSS(channelId)
// }

// happenEvery(time.Second*10, scrapeLukeJ)

// err = uploadVideoSnapshot(db, videoId)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

func main() {
	// db, err := sql.Open("sqlite3", "database.db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/channel", func(c echo.Context) error {
		channelId := c.QueryParam("channelId")
		channelSnapshot, err := getCurrentChannelSnapshot(channelId)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, channelSnapshot)
	})

	e.GET("/video", func(c echo.Context) error {
		videoId := c.QueryParam("videoId")
		videoSnapshot, err := getCurrentVideoSnapshot(videoId)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, videoSnapshot)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
