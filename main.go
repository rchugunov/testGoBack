package main

import (
	"com/github/rchugunov/awesomeProject/events"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/russross/blackfriday.v2"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(static.Serve("/", static.LocalFile("./static", false))) // static files have higher priority over dynamic routes
	router.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
	})
	api := router.Group("/api")
	{
		api.GET("/event", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.GET("/events", func(c *gin.Context) {
			events.TestEvent(c)
		})
	}

	router.Run(":" + port)
}
