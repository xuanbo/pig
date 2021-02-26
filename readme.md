# pig

## Introduction

pig is a standalone web framework

## Feature

- logger（zap）
- config（viper）
- http（echo）
- orm（gorm）

## Install

```shell
go get -u github.com/xuanbo/pig
```

## Quick Start

### Configuration

config.yaml

```yml
logger:
  level: debug
http:
  addr: :9090
datasource:
  dialect: mysql
  dns: root:123456@tcp(127.0.0.1:3306)/pig?charset=utf8&parseTime=True&loc=Local
  # 数据库连接池
  pool:
    maxIdleConns: 4
    maxOpenConns: 32
    connMaxLifetime: 1m
  sql:
    # 慢SQL
    slowThreshold: 200ms
```

### Example

```go
func main() {
	// 初始化
	pig.Initialize()

	// 注册HTTP服务
	httpServer := http.DefaultServer()
	httpServer.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})
	pig.Serve(httpServer)

	// 运行
	if err := pig.Run(); err != nil {
		t.Error(err)
		return
	}
}
```
