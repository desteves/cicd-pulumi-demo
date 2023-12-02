// Package api contains all http-related code
package api

import "github.com/gin-gonic/gin"

// endpoints
const (
	Version             = "/v1"
	HealthcheckEndpoint = "/hello"
)

// handlers
func HealthcheckHandler(c *gin.Context) {
	c.JSON(200, "world")
}

func Run() {
	// http server
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET(HealthcheckEndpoint, HealthcheckHandler)
	}

	router.Run(":8080")
}
