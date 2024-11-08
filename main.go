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

	"horizon/config"
	"horizon/core"
)

func main() {
	if _, err := os.Stat(".env.local"); err == nil {
		if err := godotenv.Load(".env.local"); err != nil {
			panic(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	config.MongoDatabase = os.Getenv("MONGO_DATABASE")

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	HORIZON_AUTH_KEY := os.Getenv("HORIZON_AUTH_KEY")
	// ALLOWED_ROUTES := []string{os.Getenv("PRIMARY_ALLOWED_ROUTE")}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: ALLOWED_ROUTES,
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Horizon is up and running!")
	})

	e.POST("/tracked-channels", func(c echo.Context) error {
		channelId := c.QueryParam("channelId")
		key := c.QueryParam("key")
		if key != HORIZON_AUTH_KEY {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		newTrackedChannel, err := core.AddTrackedChannel(channelId, mongoClient)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, newTrackedChannel)
	})

	e.GET("/tracked-channels", func(c echo.Context) error {
		trackedChannels, err := core.GetAllTrackedChannels(mongoClient)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, trackedChannels)
	})

	// e.GET("/video-snapshots", func(c echo.Context) error {
	// 	channelId := c.QueryParam("channelId")
	// 	videoSnapshots, err := getMostRecentVideoSnapshotsByChannelId(channelId, mongoClient)
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, err.Error())
	// 	}
	// 	return c.JSON(http.StatusOK, videoSnapshots)
	// })

	// e.POST("/upload-trigger", func(c echo.Context) error {
	// 	key := c.QueryParam("key")
	// 	if key != HORIZON_AUTH_KEY {
	// 		return c.String(http.StatusUnauthorized, "Unauthorized")
	// 	}
	// 	err := uploadTrigger(mongoClient)
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, err.Error())
	// 	}
	// 	return c.String(http.StatusOK, "Upload trigger successful")
	// })

	// e.POST("/video-update-trigger", func(c echo.Context) error {
	// 	key := c.QueryParam("key")
	// 	if key != HORIZON_AUTH_KEY {
	// 		return c.String(http.StatusUnauthorized, "Unauthorized")
	// 	}
	// 	err := videoUpdateTrigger(mongoClient)
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, err.Error())
	// 	}
	// 	return c.String(http.StatusOK, "Update trigger successful")
	// })

	// e.POST("channel-update-trigger", func(c echo.Context) error {
	// 	key := c.QueryParam("key")
	// 	if key != HORIZON_AUTH_KEY {
	// 		return c.String(http.StatusUnauthorized, "Unauthorized")
	// 	}
	// 	err := channelUpdateTrigger(mongoClient)
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, err.Error())
	// 	}
	// 	return c.String(http.StatusOK, "Update trigger successful")
	// })

	// e.GET("/video-rss", func(c echo.Context) error {
	// 	horizonUserId := c.QueryParam("horizonUserId")
	// 	rssFeed, err := getRecentVideoIdsWithRSS(horizonUserId, mongoClient)
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, err.Error())
	// 	}
	// 	return c.JSON(http.StatusOK, rssFeed)
	// })

	e.Logger.Fatal(e.Start(":1323"))
}
