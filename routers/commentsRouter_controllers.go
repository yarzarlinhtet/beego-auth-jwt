package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["auth-api/controllers:AuthController"] = append(beego.GlobalControllerRouter["auth-api/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "ApiLogin",
			Router:           "/login",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["auth-api/controllers:AuthController"] = append(beego.GlobalControllerRouter["auth-api/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/signup",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["auth-api/controllers:UserController"] = append(beego.GlobalControllerRouter["auth-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Me",
			Router:           "/me",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
