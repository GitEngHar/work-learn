package main

import (
	"fmt"
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/goldi/validation"
)

type Logger interface {
	Log(msg string)
}

type SimpleLogger struct{}

func (l *SimpleLogger) Log(msg string) {
	fmt.Println("[Log] ", msg)
}

type Greeter struct {
	appName string
	logger  Logger
}

func NewGreeter(appName string, logger Logger) *Greeter {
	return &Greeter{appName: appName, logger: logger}
}

func (g *Greeter) Greet(target string) {
	greeting := fmt.Sprintf("Hello, %s, Welcome to %s!", g.appName, target)
	g.logger.Log(greeting)
}

// DIコンテナの初期化
func main() {
	// TypeRegistryとパラメータマップを用意
	registry := goldi.NewTypeRegistry()
	params := map[string]interface{}{
		"app_name": "MyApp",
	}
	container := goldi.NewContainer(registry, params)
	// サービスの登録
	// 構造体インスタンスを直接登録
	container.RegisterType("logger", &SimpleLogger{})

	// ファクトリ関数を登録
	container.RegisterType("greeter", NewGreeter, "%app_name%", "@logger")

	// コンテナ検証
	validator := validation.NewContainerValidator()
	if err := validator.Validate(container); err != nil {
		panic(err)
	}

	// サービス取得&利用
	greeter := container.MustGet("greeter").(*Greeter)
	greeter.Greet("World")
}
