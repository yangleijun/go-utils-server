package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:DESController"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:DESController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/3des/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:DESController"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:DESController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/3des/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:MD5Controller"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:MD5Controller"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/md5/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:MD5Controller"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:MD5Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/md5/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Error",
            Router: `/message/error`,
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:QRCodeController"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:QRCodeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/qrcode/`,
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:QRCodeController"] = append(beego.GlobalControllerRouter["github.com/yangleijun/go-utils-server/controllers:QRCodeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/qrcode/`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
