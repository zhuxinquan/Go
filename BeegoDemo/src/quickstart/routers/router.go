package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.RouterCaseSensitive = false
    beego.Router("/abc", &controllers.MainController{})
    beego.Router("/", &controllers.AddController{})
}
