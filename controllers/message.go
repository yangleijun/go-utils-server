package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yangleijun/go-utils-server/models/views"
)

type MessageController struct {
	beego.Controller
}

// @router /message/error  [GET]
func (this *MessageController) Error() {
	model := this.GetSession("message.error")
	data := model.(views.PageError)
	this.Data["Data"] = data
	this.TplName = "message/error.tpl"
}
