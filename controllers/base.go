package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yangleijun/go-utils-server/models/views"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) errorDeal(title, msg, backUrl string) {
	errorMsg := views.PageError{}
	errorMsg.Title = title
	errorMsg.Msg = msg
	errorMsg.BackUrl = backUrl
	this.SetSession("message.error", errorMsg)
	this.Redirect("/message/error", 302)
	this.StopRun()
}
