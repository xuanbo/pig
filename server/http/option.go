package http

import (
	"fmt"
	"net/http"

	"github.com/xuanbo/pig"
	"github.com/xuanbo/pig/model"
	"github.com/xuanbo/pig/server/http/middleware"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Option 配置HTTPServer
type Option func(s *Server)

// Use 使用Echo实例
func Use(f func(*echo.Echo)) Option {
	return func(s *Server) {
		f(s.echo)
	}
}

// AddLoggerMiddleware 添加logger中间件
func AddLoggerMiddleware(message string, ctxVariables ...string) Option {
	return func(s *Server) {
		s.echo.Use(middleware.NewZapLogger(message, ctxVariables...))
	}
}

// AddRecoverMiddleware 添加recover中间件
func AddRecoverMiddleware() Option {
	return func(s *Server) {
		s.echo.Use(middleware.NewRecover())
	}
}

// AddCORSMiddleware 添加跨域中间件
func AddCORSMiddleware() Option {
	return func(s *Server) {
		s.echo.Use(middleware.NewCORS())
	}
}

// AddJWTMiddleware 添加jwt中间件
func AddJWTMiddleware() Option {
	return func(s *Server) {
		s.echo.Use(middleware.NewJWT())
	}
}

// WithAddr 配置监听地址
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithHTTPErrorHandler 自定义HTTP错误处理
func WithHTTPErrorHandler(h echo.HTTPErrorHandler) Option {
	return func(s *Server) {
		s.echo.HTTPErrorHandler = h
	}
}

// DefaultHTTPErrorHandlerFunc 默认HTTP错误处理
func DefaultHTTPErrorHandlerFunc() echo.HTTPErrorHandler {
	return func(e error, c echo.Context) {
		if he, ok := e.(*echo.HTTPError); ok {
			message := fmt.Sprintf("%s", he.Message)
			if err := c.JSON(he.Code, model.Fail(message)); err != nil {
				pig.Logger().Error("统一HTTP错误处理响应错误", zap.Error(err))
			}
			return
		}
		if err := c.JSON(http.StatusOK, model.Fail(e.Error())); err != nil {
			pig.Logger().Error("统一HTTP错误处理响应错误", zap.Error(err))
		}
	}
}
