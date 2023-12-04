package main

import (
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/api"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowFiles:       true,
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
	}))

	pgClient := db.ConnectToPostgres()

	var (
		userHandler           = api.NewUserHandler(db.NewPGUserStore(pgClient))
		projectHandler        = api.NewProjectHandler((db.NewPGProjectStore(pgClient)))
		codingActivityHandler = api.NewCodingActivityHandler(db.NewPGCodingActivityStore(pgClient))
		analyticsHandler      = api.NewAnalyticsHandler(db.NewPGAnalyticsStore(pgClient))
		languageHandler       = api.NewLanguageActivityHandler(db.NewPGLanguageStore(pgClient))
	)

	userGroup := r.Group("/user")
	{
		userGroup.POST("", userHandler.HandleCreateUser)
		userGroup.POST("/login", userHandler.HandleUserLogin)
	}

	projectGroup := r.Group("/project")
	projectGroup.Use(middleware.CheckAPIKey(pgClient))
	{
		projectGroup.POST("/add", projectHandler.HandleAddProject)
	}

	codingActivityGroup := r.Group("/coding-activity")
	codingActivityGroup.Use(middleware.CheckAPIKey(pgClient))
	{
		codingActivityGroup.POST("/", codingActivityHandler.HandleUpdateCodingActivity)
	}

	analyticsGroup := r.Group("/analytics")
	{
		analyticsGroup.GET("/coding-time", analyticsHandler.HandleGetCodingTime)
		// analyticsGroup.GET("/coding-time-range", analyticsHandler.HandleGetCodingTimeRangePerProject)
		analyticsGroup.GET("/coding-statistics", analyticsHandler.HandleGetCodingStatistics)
		analyticsGroup.GET("/language-time", analyticsHandler.HandleGetLanguageTime)
		analyticsGroup.GET("/average-time", analyticsHandler.HandleGetAverageTime)
		analyticsGroup.GET("/project-time", analyticsHandler.HandleGetProjectTime)
		analyticsGroup.GET("/leaderboard", analyticsHandler.HandleGetLeaderboard)
	}

	languageGroup := r.Group("/language-activity")
	languageGroup.Use(middleware.CheckAPIKey(pgClient))
	{
		languageGroup.POST("/", languageHandler.UpdateLanguageActivity)
	}

	r.Run()
}
