package middleware

import (
	"auth-api/service"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strings"
)

func ApiAuthFilter(ctx *context.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString != "" {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	token, err := service.NewJWTService().ValidateToken(tokenString)

	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Unauthorized"))
	} else if !token.Valid {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Unauthorized"))
	}
}
