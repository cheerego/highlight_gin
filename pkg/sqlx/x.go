package sqlx

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	Id                 int
	First, Last, Email string
}

var Sqlx *sqlx.DB

func init() {
	var err error
	Sqlx, err = sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	//user := []User{}
	//
	//db.Select(&user, "select * from users")
	//
	//log.Println("users...")
	//log.Println(user)
}