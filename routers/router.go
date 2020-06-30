// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"anydevelop.cn/recruit-server/controllers"
	"anydevelop.cn/recruit-server/filter"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, filter.CORS())
	beego.Include(&controllers.LoginAdmin{})
	admin := beego.NewNamespace("/Admin", beego.NSInclude(&controllers.AdminController{}))
	beego.AddNamespace(admin)
}
