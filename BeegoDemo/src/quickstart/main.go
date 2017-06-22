package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func main() {
	beego.Include(&controllers.CMSController{})
	beego.Run()
}