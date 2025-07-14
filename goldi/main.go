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
	registry.RegisterAll(map[string]goldi.TypeFactory{
		"user.domain": goldi.NewType(
			NewUser,
			"%name%",
			"%message%",
		),
		"user.domain.usecase": goldi.NewType(
			NewCreateUserUsecase,
			"@user.domain",
		),
		"user.domain.handler": goldi.NewType(
			CreateUserHandler,
			"@user.domain.usecase",
		),
	})

	// ハンドラを登録

	container := goldi.NewContainer(registry, params)
	// サービス取得&利用
	userResponse := container.MustGet("user.domain.handler").(*Response)
	fmt.Println(*userResponse)
}
