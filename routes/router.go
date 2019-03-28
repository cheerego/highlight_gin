package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/db", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"result": "123",
		})
	})

	return r
}
