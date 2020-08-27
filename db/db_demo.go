package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

var _ error
var engin *gorose.Engin

func init() {
	engin, _ = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true"})
}
func DB() gorose.IOrm {
	return engin.NewOrm()
}

type Test struct {
	Age  int64 `json:"age"`  //
	Time int64 `json:"time"` //
}

func (t *Test) TableName() string {
	return "test"
}

func main() {
	db := DB()

	t := Test{
		Age:  15,
		Time: 15,
	}
	db.Table(&t).Data(t).Where("age", "=", 14).Update()
	fmt.Println(t)
}
