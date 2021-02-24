package pig_test

import (
	"testing"

	"github.com/xuanbo/pig"
	"github.com/xuanbo/pig/server/http"

	"github.com/labstack/echo/v4"
)

func TestApp(t *testing.T) {
	// 初始化
	pig.Initialize()

	// 注册HTTP服务
	httpServer := http.DefaultServer()
	httpServer.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})
	pig.Serve(httpServer)

	// 运行
	pig.Run()
}
