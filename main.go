package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", homePage)

	router.Run(":" + "80")
}

func homePage(c *gin.Context) {
	c.String(200, "Welcome to Go test page")
}
