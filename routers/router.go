package routers

import (
	"github.com/astaxie/beego"
	"github.com/yangleijun/go-utils-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.QRCodeController{})
	beego.Include(&controllers.MessageController{})
	beego.Include(&controllers.MD5Controller{})
	beego.Include(&controllers.DESController{})
}
