package pig

import (
	"fmt"
	"sync"

	"github.com/xuanbo/pig/config"
	"github.com/xuanbo/pig/logger"
	"github.com/xuanbo/pig/server"
)

var (
	initOnce sync.Once
	app      *App
)

// App 程序
type App struct {
	config  *config.Config
	logger  *logger.Logger
	servers []server.Server

	hideBanner bool
	banner     string
}

// Initialize 初始化
func Initialize(options ...Option) {
	initOnce.Do(func() {
		app = initialize(options...)
		printBanner()
	})
}

func printBanner() {
	if app.hideBanner {
		return
	}
	fmt.Println(app.banner)
}

// initialize 初始化
func initialize(options ...Option) *App {
	app := &App{
		config:     config.New(),
		servers:    make([]server.Server, 0, 4),
		hideBanner: false,
		banner: `
        (_) 
  _ __  _  __ _ 
  |  _  | |  _  |     author: xuanbo.wang
  | |_) | | (_| |     email: 1345545983@qq.com
  | .__/|_|\__, |     github: https://github.com/xuanbo
  | |       __/ |
  |_|      |___/ 
`,
	}
	level := app.config.GetString("logger.level")
	switch level {
	case "debug":
		app.logger = logger.New(logger.DebugLevel)
	case "info":
		app.logger = logger.New(logger.InfoLevel)
	case "warn":
		app.logger = logger.New(logger.WarnLevel)
	case "error":
		app.logger = logger.New(logger.ErrorLevel)
	default:
		panic("配置 [logger.level] 值不合法，请检查配置")
	}
	for _, option := range options {
		option(app)
	}
	return app
}

// Serve 注册HTTP服务
func Serve(server server.Server) {
	Initialize()
	app.servers = append(app.servers, server)
}

// Run 运行
func Run() error {
	Initialize()
	defer app.logger.Flush()
	for _, server := range app.servers {
		if err := server.Serve(); err != nil {
			return err
		}
	}
	return nil
}

// Config 配置实例
func Config() *config.Config {
	Initialize()
	return app.config
}

// Logger logger实例
func Logger() *logger.Logger {
	Initialize()
	return app.logger
}
