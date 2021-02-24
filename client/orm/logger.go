package orm

import (
	"context"
	"fmt"
	"time"

	"github.com/xuanbo/pig"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// Logger 实现日志
type Logger struct {
	SlowThreshold time.Duration
	Level         logger.LogLevel
	Message       string
}

// LogMode 实现LogMode接口
func (lg *Logger) LogMode(level logger.LogLevel) logger.Interface {
	lg.Level = level
	return lg
}

// Info 实现Info接口
func (lg *Logger) Info(ctx context.Context, msg string, v ...interface{}) {
	pig.Logger().Debug(fmt.Sprintf(msg, v...))
}

// Warn 实现Warn接口
func (lg *Logger) Warn(ctx context.Context, msg string, v ...interface{}) {
	pig.Logger().Warn(fmt.Sprintf(msg, v...))
}

// Error 实现Error接口
func (lg *Logger) Error(ctx context.Context, msg string, v ...interface{}) {
	pig.Logger().Error(fmt.Sprintf(msg, v...))
}

// Trace 实现Trace接口
func (lg *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if lg.Level > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && lg.Level >= logger.Error:
			sql, rows := fc()
			pig.Logger().Error(lg.Message, zap.String("sql", sql), zap.String("cost", elapsed.String()), zap.Int64("rows", rows), zap.Error(err))
		case lg.Level >= logger.Info:
			sql, rows := fc()
			if elapsed >= lg.SlowThreshold {
				pig.Logger().Warn(lg.Message, zap.String("sql", sql), zap.String("cost", elapsed.String()), zap.Int64("rows", rows))
			} else {
				pig.Logger().Debug(lg.Message, zap.String("sql", sql), zap.String("cost", elapsed.String()), zap.Int64("rows", rows))
			}
		}
	}
}
