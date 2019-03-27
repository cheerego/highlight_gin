package routers

import (
	_ "git.catchme.cn/placeless/highlight_gin/pkg/dotenv"
	_ "git.catchme.cn/placeless/highlight_gin/pkg/logging"
	"git.catchme.cn/placeless/highlight_gin/pkg/sqlx"
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
			"result": sqlx.Sqlx.QueryRow("select * from users"),
		})
	})

	return r
}
