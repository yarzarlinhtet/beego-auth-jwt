package controllers

import (
	"auth-api/dto"
	"auth-api/models"
	"auth-api/service"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) ApiResponse(code int, data map[string]interface{}) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = data
	c.ServeJSON()
}

// @Title SignUp
// @Description create users
// @Param	body  body  models.User  true  "The object content"
// @Failure 403 body is empty
// @router /me [get]
func (c *UserController) Me() {
	token := c.Ctx.Request.Header.Get("Authorization")
	fmt.Println("bearer token : " + token)
	if token != "" {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	fmt.Println("token : " + token)

	meta, err := service.NewJWTService().Meta(token)
	if err != nil {
		c.ApiResponse(401, map[string]interface{}{"data": "", "message": "Invalid token"})
	}
	user := models.User{Username: meta.Name}
	user.Read("Username")
	c.ApiResponse(200, map[string]interface{}{"data": dto.UserResponse{Username: user.Username, Email: user.Email}, "message": "success"})
}
