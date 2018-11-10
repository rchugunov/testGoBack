package main

import (
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

	router.Run(":" + port)
}
