package main

import "github.com/fgrosse/goldi"

func InitializeTesting(f func(types goldi.TypeRegistry)) *goldi.Container {
	registry := goldi.NewTypeRegistry()
	RegisterType(registry)
	f(registry)
	//params := envMap()
	params := map[string]interface{}{
		"name":    "userA",
		"message": "hello",
	}
	c := goldi.NewContainer(registry, params)
	return c
}
