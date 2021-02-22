package middleware

import (
	"net/http"
	"time"

	"github.com/xuanbo/pig"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger zap日志中间件
func NewZapLogger(message string, ctxVariables ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			req := ctx.Request()
			resp := ctx.Response()
			start := time.Now()
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}
			cost := time.Since(start)

			// 字段
			fields := []zapcore.Field{
				zap.String("uri", req.RequestURI), zap.String("method", req.Method),
				zap.Int("status", resp.Status), zap.String("cost", cost.String()),
			}
			for _, name := range ctxVariables {
				fields = append(fields, zap.Any(name, ctx.Get(name)))
			}

			if err == nil || resp.Status == http.StatusNotFound {
				pig.Logger().Debug(message, fields...)
			} else {
				fields = append(fields, zap.Error(err))
				pig.Logger().Warn(message, fields...)
			}
			return nil
		}
	}
}
