package main

import (
	"fmt"
	_ "git.catchme.cn/placeless/highlight_gin/pkg/db"
	_ "git.catchme.cn/placeless/highlight_gin/pkg/dotenv"
	"git.catchme.cn/placeless/highlight_gin/routes"
	"gopkg.in/guregu/null.v3"
)

type Student struct {
	Id       int
	Name     null.String
	NickName null.Int
}

func main() {
	var s = Student{}
	name, _ := s.Name.Value()
	value, _ := s.NickName.Value()
	fmt.Print(s.Id, ":", name, ":", value)
	r := routers.InitRouter()
	//gin.SetMode(gin.DebugMode)
	r.Run() // listen and serve on 0.0.0.0:8080
}
