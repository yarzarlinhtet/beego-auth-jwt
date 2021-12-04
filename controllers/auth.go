package controllers

import (
	"auth-api/dto"
	"auth-api/models"
	"auth-api/service"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) ApiResponse(code int, data map[string]interface{}) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = data
	c.ServeJSON()
}

// @Title SignUp
// @Description create users
// @Param	body  body  models.User  true  "The object content"
// @Failure 403 body is empty
// @router /signup [post]
func (c *AuthController) Post() {
	user := models.User{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	if err != nil {
		fmt.Println("error : " + err.Error())
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	fmt.Println("hash : " + string(hashedPassword))
	user.Password = string(hashedPassword)
	err = user.Insert()
	if err != nil {
		fmt.Println("error : " + err.Error())
	}
	c.Data["json"] = user
	c.ServeJSON()
}

// @Title Login
// @Description login user
// @Param	body  body  dto.LoginCredential  true  "The object content"
// @Failure 403 body is empty
// @router /login [post]
func (c *AuthController) ApiLogin() {
	var credential dto.LoginCredential
	user := models.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &credential)
	if err != nil {
		fmt.Println("testing : " + err.Error())
	}
	o := orm.NewOrm()

	user.Username = credential.Username

	err = o.Read(&user, "Username")

	if err != nil {
		//c.Ctx.Output.SetStatus(422)
		//c.Data["json"] = map[string]string{"message": "Username or password invalid!"}
		//c.ServeJSON()
		c.ApiResponse(422, map[string]interface{}{"data": "", "message": "Username or password invalid!"})
	}

	var jwtService service.JWTService = service.NewJWTService()
	var loginService service.LoginService = service.NewLoginService(user)
	var authService service.AuthService = service.NewAuthService(loginService, jwtService)

	if isAuth := authService.Login(credential.Username, credential.Password); !isAuth {
		c.ApiResponse(422, map[string]interface{}{"data": "", "message": "Username or password invalid!"})

	}

	c.ApiResponse(422, map[string]interface{}{"data": map[string]interface{}{"apikey": jwtService.GenerateToken(user.Username, true)}, "message": "Username or password invalid!"})

}
