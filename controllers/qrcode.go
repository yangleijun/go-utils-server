package controllers

import (
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
	"github.com/yangleijun/go-utils-server/models/views"
)

type QRCodeController struct {
	beego.Controller
}

func (this *QRCodeController) errorDeal(errorMsg *views.PageError) {
	this.SetSession("message.error", *errorMsg)
	this.Redirect("/message/error", 302)
}

// @router /qrcode/  [GET]
func (c *QRCodeController) Get() {
	c.TplName = "qrcode/get.tpl"
}

// @router /qrcode/  [POST]
func (this *QRCodeController) Post() {
	url := this.GetString("url")
	if len(url) == 0 {
		errorMsg := views.PageError{}
		errorMsg.Title = "参数错误"
		errorMsg.Msg = "未传递参数url,请输入输入需要生成的二维码信息"
		errorMsg.BackUrl = "/qrcode/"
		this.errorDeal(&errorMsg)
		return
	}
	images, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		errorMsg := views.PageError{}
		errorMsg.Title = "二维码生成错误!"
		errorMsg.Msg = err.Error()
		errorMsg.BackUrl = "/qrcode/"
		return
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "image/png")
	this.Ctx.ResponseWriter.Write(images)
}
