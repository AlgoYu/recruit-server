package controllers

import (
	"github.com/astaxie/beego"
	"recruit-server/common"
)

type AdminBase struct {
	beego.Controller
}

// 从请求头中获取并验证AdminToken
func (adminBase *AdminBase) Prepare() {
	adminTokenStr := adminBase.Ctx.Request.Header.Get("Admin-Token")
	if adminTokenStr == "" {
		adminBase.Data["json"] = common.Fail("No token.")
		adminBase.ServeJSON()
		adminBase.StopRun()
	}
}
