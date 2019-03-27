package dotenv

import (
	"fmt"
	"git.catchme.cn/placeless/highlight_gin/pkg/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var db *gorm.DB

//type Model struct {
//	ID         int `gorm:"primary_key" json:"id"`
//	CreatedOn  int `json:"created_on"`
//	ModifiedOn int `json:"modified_on"`
//	DeletedOn  int `json:"deleted_on"`
//}

// Setup initializes the database instance
func init() {
	var err error
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	))
	//
	if err != nil {
		logging.Logger.Error("数据库连接失败！")
		os.Exit(0)
	}
	//
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return setting.DatabaseSetting.TablePrefix + defaultTableName
	//}

	//db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
//func updateTimeStampForCreateCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		nowTime := time.Now().Unix()
//		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
//			if createTimeField.IsBlank {
//				createTimeField.Set(nowTime)
//			}
//		}
//
//		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
//			if modifyTimeField.IsBlank {
//				modifyTimeField.Set(nowTime)
//			}
//		}
//	}
//}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
//func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		scope.SetColumn("ModifiedOn", time.Now().Unix())
//	}
//}

// deleteCallback will set `DeletedOn` where deleting
//func deleteCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		var extraOption string
//		if str, ok := scope.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
//
//		if !scope.Search.Unscoped && hasDeletedOnField {
//			scope.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v%v%v",
//				scope.QuotedTableName(),
//				scope.Quote(deletedOnField.DBName),
//				scope.AddToVars(time.Now().Unix()),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		} else {
//			scope.Raw(fmt.Sprintf(
//				"DELETE FROM %v%v%v",
//				scope.QuotedTableName(),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		}
//	}
//}

// addExtraSpaceIfExist adds a separator
//func addExtraSpaceIfExist(str string) string {
//	if str != "" {
//		return " " + str
//	}
//	return ""
//}
