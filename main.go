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

func main() {

	if err := godotenv.Load(".env.local"); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	HORIZON_AUTH_KEY := os.Getenv("HORIZON_AUTH_KEY")
	ALLOWED_ROUTES := []string{os.Getenv("PRIMARY_ALLOWED_ROUTE")}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: ALLOWED_ROUTES,
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Horizon is up and running!")
	})

	e.POST("/tracked-channel", func(c echo.Context) error {
		channelId := c.QueryParam("channelId")
		key := c.QueryParam("key")
		if key != HORIZON_AUTH_KEY {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		newTrackedChannel, err := addTrackedChannel(channelId, mongoClient)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, newTrackedChannel)
	})

	e.GET("/tracked-channel", func(c echo.Context) error {
		trackedChannels, err := getAllTrackedChannels(mongoClient)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, trackedChannels)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
