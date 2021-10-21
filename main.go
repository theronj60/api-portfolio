package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theronj60/api-portfolio/app"
)

func main () {
	app.LoadEnv()
	app.Connect()

	router := gin.Default()
	router.GET("/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})
	router.Run(":8080")
}
