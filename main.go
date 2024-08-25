package main

import (
	"database/sql"
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

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/channel", func(c echo.Context) error {
		channelId := c.FormValue("channelId")
		err = createChannelFromId(db, channelId)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "Channel created")
	})

	e.Logger.Fatal(e.Start(":1323"))
	// happenEvery(time.Second*10, scrapeLukeJ)

	// videoId := "3GAAIKDUMLg"

	// err = uploadVideoSnapshot(db, videoId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
