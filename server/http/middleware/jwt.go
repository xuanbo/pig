package middleware

import (
	"fmt"
	"net/http"

	"github.com/xuanbo/pig"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewJWT jwt中间件
func NewJWT() echo.MiddlewareFunc {
	var (
		secret    = pig.Config().GetString("http.jwt.secret")
		skipPaths = pig.Config().GetStringSlice("http.jwt.skipPaths")
	)
	if secret == "" {
		secret = "secret"
	}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		// 跳过登录
		Skipper: func(ctx echo.Context) bool {
			for _, skipPath := range skipPaths {
				if ctx.Path() == skipPath {
					return true
				}
			}
			return false
		},
		SigningKey:  []byte(secret),
		ContextKey:  "JWT_TOKEN",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		ErrorHandler: func(err error) error {
			if he, ok := err.(*echo.HTTPError); ok {
				message := fmt.Sprintf("%s", he.Message)
				return echo.NewHTTPError(http.StatusUnauthorized, message)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		},
		SuccessHandler: func(ctx echo.Context) {
			if token, ok := ctx.Get("JWT_TOKEN").(*jwt.Token); ok {
				userID := token.Header["userId"]
				userName := token.Header["userName"]
				username := token.Header["username"]
				ctx.Set("JWTuserID", userID)
				ctx.Set("JWTUserName", userName)
				ctx.Set("JWTUsername", username)
			}
		},
	})
}
