package controllers

import (
	"github.com/astaxie/beego"
)

// CMS API
type CMSController struct {
	beego.Controller
}

//func (c *CMSController) URLMapping() {
//	c.Mapping("StaticBlock", c.StaticBlock)
//	c.Mapping("AllBlock", c.AllBlock)
//}


// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":key")))
}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {

}
