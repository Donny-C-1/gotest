package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", homePage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router.Run(":" + "80")
}

func homePage(c *gin.Context) {
	c.String(200, "Welcome to Go test page")
}
