package dotenv

import (
	"git.catchme.cn/placeless/highlight_gin/pkg/logging"
	"github.com/joho/godotenv"
	"os"
	"path"
)

func init() {
	_ = godotenv.Load()
	pwd, _ := os.Getwd()
	var err error = godotenv.Load(path.Join(pwd, ".env"))
	if err != nil {
		logging.Logger.Error("配置读取失败")
		os.Exit(0)
	}

	logging.Logger.Info("配置读取成功", os.Getenv("PATH"))
}
