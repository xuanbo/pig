# pig

## Introduction

pig is a standalone web framework

## Feature

- logger（zap）
- config（viper）
- http（echo）

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
