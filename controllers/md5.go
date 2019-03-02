package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type MD5Controller struct {
	BaseController
}

// @router /md5/ [get]
func (this *MD5Controller) Get() {
	this.TplName = "md5/get.tpl"
}

// @router /md5/ [post]
func (this *MD5Controller) Post() {
	values := this.GetString("values")
	if len(values) == 0 {
		this.errorDeal("参数错误", "未传递values的值", "/md5/")
	}
	m := md5.New()
	m.Write([]byte(values))
	resultBytes := m.Sum(nil)
	fmt.Println(resultBytes)
	resultStr := hex.EncodeToString(resultBytes)
	this.Ctx.WriteString(resultStr)
}
