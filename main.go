package main

import (
	"fmt"

	"os"
	"github.com/gin-gonic/gin"
	_ "git.catchme.cn/placeless/highlight_gin/pkg/dotenv"
)

func main() {

	fmt.Print(os.Getenv("asd"))
	r := gin.Default()
	//gin.SetMode(gin.DebugMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
