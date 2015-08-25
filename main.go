package main
import (
	"github.com/astaxie/beego"
	"github.com/zihuxinyu/youzan/sdk"
)

func main() {
	beego.Error("main")
	clt := sdk.NewClient("ss", "ssss", nil)
	type Tmp struct {
		Method string
		Zhang string
		San   int
		Li    string
		Si    int64
	}

	tmp := new(Tmp)
	tmp.Zhang = "zhang"
	tmp.San = 3
	tmp.Li = "li"
	tmp.Si = 64
	tmp.Method=""

	clt.PostJSON( tmp, nil)
}
