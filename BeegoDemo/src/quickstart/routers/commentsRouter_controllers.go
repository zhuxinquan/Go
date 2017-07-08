package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["quickstart/controllers:CMSController"] = append(beego.GlobalControllerRouter["quickstart/controllers:CMSController"],
		beego.ControllerComments{
			Method: "StaticBlock",
			Router: `/staticblock/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["quickstart/controllers:CMSController"] = append(beego.GlobalControllerRouter["quickstart/controllers:CMSController"],
		beego.ControllerComments{
			Method: "AllBlock",
			Router: `/all/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
