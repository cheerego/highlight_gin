package main

import (
	_ "git.catchme.cn/placeless/highlight_gin/pkg/dotenv"
	"git.catchme.cn/placeless/highlight_gin/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := routers.InitRouter()
	//gin.SetMode(gin.DebugMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": os.Getenv("PORT"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
