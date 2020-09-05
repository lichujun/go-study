package beego_demo

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModelWithPrefix("t_", new(User))
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func TestQuery(t *testing.T) {
	o := orm.NewOrm()

	o.QueryTable("t_user").Filter("name", "joseph").Filter("age__gt", "12").Update(orm.Params{
		"name": "chandler",
	})

	var users []User
	//o.QueryTable("t_user").All(&users)

	builder, _ := orm.NewQueryBuilder("mysql")
	builder.Select("name", "age").
		From("t_user")
	o.Raw(builder.String()).QueryRows(&users)

	fmt.Println(users)
}
