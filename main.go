package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	Name  string `gorm:"primaryKey;column:name"`
	Phone string `gorm:"primaryKey;column:phone"`
	Age   int    `gorm:"primaryKey;column:age"`
	Msg   string `gorm:"column:msg"`
}

func main() {
	_ = DB.AutoMigrate(&user{})

	user := user{
		Name:  "",
		Phone: "",
		Age:   22,
		Msg:   "bbb",
	}
	DB.Select([]string{"name", "phone", "age", "msg"}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "phone"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "phone", "age", "msg"}),
	}).Create(&user)
}

// DB 对象
var DB *gorm.DB

func init() {
	var err error
	// DB, err = NewSQLiteDB("arknights.db")
	DB, err = NewMySQL("127.0.0.1", "root", "", "test")
	if err != nil {
		fmt.Println("数据库初始化异常", err.Error())
	}

}

func NewMySQL(host, user, pass, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, pass, host, dbname)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
}
