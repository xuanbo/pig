package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xuanbo/pig"
)

// NewCORS 跨域中间件
func NewCORS() echo.MiddlewareFunc {
	var (
		allowOrigins     = pig.Config().GetStringSlice("http.cors.allowOrigins")
		allowHeaders     = pig.Config().GetStringSlice("http.cors.allowHeaders")
		allowMethods     = pig.Config().GetStringSlice("http.cors.allowMethods")
		allowCredentials = pig.Config().GetBool("http.cors.allowCredentials")
		maxAge           = pig.Config().GetInt("http.cors.maxAge")
	)
	if allowOrigins == nil {
		allowOrigins = []string{"*"}
	}
	if allowHeaders == nil {
		allowHeaders = []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}
	}
	if allowMethods == nil {
		allowMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}
	}
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     allowHeaders,
		AllowMethods:     allowMethods,
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge,
	})
}
