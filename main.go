package main

import (
	"fmt"
	"git.catchme.cn/placeless/highlight_gin/plugins"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	plugins.Logger.Error("123")
	plugins.Logger.Info("123")
	fmt.Print(os.Getenv("haha"))
	r := gin.Default()
	//gin.SetMode(gin.DebugMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
