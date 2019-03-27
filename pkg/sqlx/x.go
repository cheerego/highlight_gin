package sqlx

import (
	"git.catchme.cn/placeless/highlight_gin/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

func init() {
	var err error
	Sqlx, err = sqlx.Connect("mysql", "root:root@(localhost:3306)/highlight_admin")
	if err != nil {
		logging.Logger.ExitFunc = func(i int) {
			logging.Logger.Error("数据库连接失败！")
		}
		logging.Logger.Exit(0)
		os.Exit(0)
	}

	err = Sqlx.Ping()
	if err != nil {
		logging.Logger.Error("数据库连接失败！")
		os.Exit(0)
	}
}

type User struct {
	Id             int
	Name, NikeName string
}

var Sqlx *sqlx.DB
