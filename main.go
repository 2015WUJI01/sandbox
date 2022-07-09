package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Data struct {
	ID      int64  `json:"id" orm:"pk;auto;column(id)"`
	Name    string `json:"name" orm:"column(name)"`
	Content string `json:"content" orm:"column(content);type(text)"`
}

func (d Data) TableName() string {
	return "beego_multi_insert"
}

func init() {
	orm.RegisterModel(new(Data))
}

func main() {
	orm.Debug = true
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8mb4")

	sqldb, _ := orm.GetDB()
	start := time.Now()
	for i := 0; i < 1; i++ {
		// query, data := ten()
		// _, _ = o.Raw(query, data...).Exec()
		res, err := sqldb.Exec("INSERT INTO `beego_multi_insert` (`name`, `content`) VALUES (?, ?);"+
			"INSERT INTO `beego_multi_insert` (`name`, `content`) VALUES (?, ?);",
			"asdf", "asdfas", "adfafd", "asdfasf",
		)
		if err != nil {
			fmt.Println(err)
		} else {
			rows, _ := res.RowsAffected()
			fmt.Println("success:", rows)
		}
	}
	cost := time.Now().Sub(start)
	fmt.Println(cost)
}

func ten() (query string, data []interface{}) {
	d := Data{Name: "test", Content: "abcdefghijklmnopqrstuvwxyz"}
	for i := 0; i < 10; i++ {
		query += "INSERT INTO `beego_multi_insert` (`name`, `content`) VALUES (?, ?);"
		data = append(data, d.Name, d.Content)
	}
	return
}
