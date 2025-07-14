package main

import (
	"fmt"
	"github.com/fgrosse/goldi"
)

// DIコンテナの初期化
func main() {
	// TypeRegistryとパラメータマップを用意
	params := map[string]interface{}{
		"name":    "userA",
		"message": "hello",
	}

	registry := goldi.NewTypeRegistry()

	// サービスの登録
	// 構造体インスタンスを直接登録
	RegisterType(registry)

	// ハンドラを登録
	container := goldi.NewContainer(registry, params)
	// サービス取得&利用
	userResponse := container.MustGet("user.domain.handler").(*Response)
	fmt.Println(*userResponse)
}
