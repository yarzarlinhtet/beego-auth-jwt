package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body  body  dto.LoginCredential  true  "The object content"
// @Failure 403 body is empty
// @router / [post]
//func (l *LoginController) ApiLogin() {
//	var credential dto.LoginCredential
//	err := json.Unmarshal(l.Ctx.Input.RequestBody, &credential)
//	if err != nil {
//		fmt.Println("testing : " + err.Error())
//	}
//	l.Data["json"] = credential
//	l.ServeJSON()
//}
