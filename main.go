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
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8mb4")

	o := orm.NewOrm()
	p, _ := o.Raw("INSERT INTO `test`.`beego_multi_insert` (`name`, `content`) VALUES (?, ?)").Prepare()
	data := Data{Name: "test", Content: "abcdefghijklmnopqrstuvwxyz"}
	start := time.Now()
	for i := 0; i < 1e5; i++ {
		// _, _ = o.Raw("INSERT INTO `beego_multi_insert` (`name`, `content`) VALUES (?, ?)", data.Name, data.Content).Exec()
		_, _ = p.Exec(data.Name, data.Content)
	}
	cost := time.Now().Sub(start)
	fmt.Println(cost)
}
