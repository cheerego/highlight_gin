package main

import (
	_ "git.catchme.cn/placeless/highlight_gin/pkg/db"
	_ "git.catchme.cn/placeless/highlight_gin/pkg/dotenv"
	"git.catchme.cn/placeless/highlight_gin/routes"
)

func main() {
	r := routers.InitRouter()
	//gin.SetMode(gin.DebugMode)
	r.Run() // listen and serve on 0.0.0.0:8080
}
