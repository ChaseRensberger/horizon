package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func scrapeLukeJ() {
// 	channelId := "UCYzV77unbAR8KiIoSm4zdUw"
// 	getRSS(channelId)
// }

// happenEvery(time.Second*10, scrapeLukeJ)

var usingFallback = false

func main() {

	if err := godotenv.Load(".env.local"); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                        // Allows all origins
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE}, // Specify allowed methods
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/channel", func(c echo.Context) error {
		channelId := c.QueryParam("channelId")
		channelSnapshot, err := getCurrentChannelSnapshot(channelId, usingFallback)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, channelSnapshot)
	})

	e.GET("/video", func(c echo.Context) error {
		videoId := c.QueryParam("videoId")
		videoSnapshot, err := getCurrentVideoSnapshot(videoId, usingFallback)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, videoSnapshot)
	})

	e.POST("/tracked-channel", func(c echo.Context) error {
		channelId := c.QueryParam("channelId")
		newTrackedChannel, err := addTrackedChannel(channelId, client, usingFallback)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, newTrackedChannel)
	})

	e.GET("/tracked-channel", func(c echo.Context) error {
		trackedChannels, err := getAllTrackedChannels(client)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, trackedChannels)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
