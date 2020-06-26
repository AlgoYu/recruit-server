package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["recruit-server/controllers:AdminController"] = append(beego.GlobalControllerRouter["recruit-server/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "AddAdmin",
			Router:           `/AddAdmin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["recruit-server/controllers:AdminController"] = append(beego.GlobalControllerRouter["recruit-server/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/GetAdmin/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["recruit-server/controllers:AdminController"] = append(beego.GlobalControllerRouter["recruit-server/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/GetAll`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["recruit-server/controllers:AdminController"] = append(beego.GlobalControllerRouter["recruit-server/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "Modify",
			Router:           `/ModifyAdmin/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["recruit-server/controllers:AdminController"] = append(beego.GlobalControllerRouter["recruit-server/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/delete/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
