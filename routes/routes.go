// will hold routes
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @TODO need to create routes
func employee () {
	router := gin.Default()
	// router.GET("/employee", controllers)
	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	router.GET("/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})
}
