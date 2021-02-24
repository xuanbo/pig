package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level 日志级别
type Level uint8

const (
	// DebugLevel DEBUG级别
	DebugLevel Level = iota
	// InfoLevel INFO级别
	InfoLevel
	// WarnLevel WARN级别
	WarnLevel
	// ErrorLevel ERROR级别
	ErrorLevel
)

// Logger 日志
type Logger struct {
	*zap.Logger
}

// New 创建日志
func New(level Level) *Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		// warn级别以上不显示堆栈
		// StacktraceKey: "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder, // 大写编码器
		// 时间
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 日志级别
	var atom zap.AtomicLevel
	switch level {
	case DebugLevel:
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	case InfoLevel:
		atom = zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	config := zap.Config{
		Level:            atom,               // 日志级别
		Development:      true,               // 开发模式，堆栈跟踪
		Encoding:         "console",          // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,      // 编码器配置
		OutputPaths:      []string{"stdout"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return &Logger{zapLogger}
}
