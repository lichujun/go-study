package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

var err error
var engin *gorose.Engin

func init() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	// mysql Dsn示例 "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true"
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true"})
}
func DB() gorose.IOrm {
	return engin.NewOrm()
}
func main() {
	// 这里定义一个变量db, 是为了复用db对象, 可以在最后使用 db.LastSql() 获取最后执行的sql
	// 如果不复用 db, 而是直接使用 DB(), 则会新建一个orm对象, 每一次都是全新的对象
	// 所以复用 db, 一定要在当前会话周期内
	db := DB()
	res, err := db.Query("select version()")
	fmt.Println(res, err)
}
