package main

import "github.com/fgrosse/goldi"

func RegisterType(typ goldi.TypeRegistry) {
	typ.RegisterAll(map[string]goldi.TypeFactory{
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
}
