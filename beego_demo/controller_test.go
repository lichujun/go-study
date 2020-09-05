package beego_demo

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"testing"
)

func init() {
	_ = logs.SetLogger(logs.AdapterConsole)
	logs.SetLogFuncCall(true)
	beego.BConfig.Listen.HTTPPort = 9000
	beego.BConfig.Listen.EnableAdmin = true

	tk := toolbox.NewTask("tk", "5 * * * * *", func() error { fmt.Println("tk"); return nil })
	toolbox.AddTask("tk", tk)
}

func TestRun(t *testing.T) {
	beego.AutoRouter(&TestController{})
	beego.Run()

}

type TestController struct {
	beego.Controller
}

func (tc *TestController) Hello() {
	tc.Ctx.WriteString("hello")
}
