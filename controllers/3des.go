package controllers

import (
	"encoding/base64"

	"github.com/yangleijun/go-utils-server/utils"
)

type DESController struct {
	BaseController
}

//@router /3des/ [get]
func (des *DESController) Get() {
	des.TplName = "3des/get.tpl"
}

//@router /3des/ [post]
func (des *DESController) Post() {
	values := des.GetString("values")
	key := des.GetString("key")

	model := utils.DesEcb{
		Values: []byte(values),
		Key:    []byte(key),
	}
	if len(model.Key) != 24 {
		des.errorDeal("参数错误", "密钥的长度必须为24", "/3des/")
	}

	if len(model.Values) == 0 {
		des.errorDeal("参数错误", "需要加密或者解密的字符串必须不为空", "/3des/")
	}
	var t int
	if t, err := des.GetInt("type"); err != nil || (t != 0 && t != 1) {
		des.errorDeal("参数错误", "未选择加密方式", "/3des/")
	}
	var result []byte
	if t == 0 {
		result, _ = model.Encrypt()
	} else {
		result, _ = model.Decrypt()
	}
	strResult := base64.StdEncoding.EncodeToString(result)
	des.Ctx.WriteString(strResult)
}
