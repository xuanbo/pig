package http

import (
	"context"

	"github.com/xuanbo/pig"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Server HTTP服务
type Server struct {
	echo *echo.Echo
	addr string
}

// DefaultServer 创建默认HTTP服务
func DefaultServer() *Server {
	return NewServer(
		// 中间件
		AddLoggerMiddleware("API"),
		AddRecoverMiddleware(),
		// 默认HTTP错误处理
		WithHTTPErrorHandler(DefaultHTTPErrorHandlerFunc()),
	)
}

// NewServer 创建HTTP服务
func NewServer(options ...Option) *Server {
	addr := pig.Config().GetString("http.addr")
	if addr == "" {
		addr = ":9090"
	}
	s := &Server{
		echo: echo.New(),
		addr: addr,
	}
	s.echo.HideBanner = true
	s.echo.HidePort = true
	for _, option := range options {
		option(s)
	}
	return s
}

// Serve 启动服务
func (s *Server) Serve() error {
	for _, route := range s.echo.Routes() {
		pig.Logger().Info("注册HTTP API路由", zap.String("method", route.Method), zap.String("name", route.Name), zap.String("path", route.Path))
	}
	pig.Logger().Info("启动HTTP服务", zap.String("addr", s.addr))
	if err := s.echo.Start(s.addr); err != nil {
		return err
	}
	return nil
}

// Stop 关闭服务
func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

// GET 注册GET请求
func (s *Server) GET(path string, handleFunc echo.HandlerFunc, mwFuncs ...echo.MiddlewareFunc) {
	s.echo.GET(path, handleFunc, mwFuncs...)
}

// POST 注册POST请求
func (s *Server) POST(path string, handleFunc echo.HandlerFunc, mwFuncs ...echo.MiddlewareFunc) {
	s.echo.POST(path, handleFunc, mwFuncs...)
}

// PUT 注册PUT请求
func (s *Server) PUT(path string, handleFunc echo.HandlerFunc, mwFuncs ...echo.MiddlewareFunc) {
	s.echo.PUT(path, handleFunc, mwFuncs...)
}

// DELETE 注册DELETE请求
func (s *Server) DELETE(path string, handleFunc echo.HandlerFunc, mwFuncs ...echo.MiddlewareFunc) {
	s.echo.DELETE(path, handleFunc, mwFuncs...)
}

// Use 注册中间件
func (s *Server) Use(mwFuncs ...echo.MiddlewareFunc) {
	s.echo.Use(mwFuncs...)
}
