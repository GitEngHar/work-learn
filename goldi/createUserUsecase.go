package main

import (
	"errors"
)

type CreateUserUsecase interface {
	Do() (string, error)
}

type createUserUsecase struct {
	user *User
}

func NewCreateUserUsecase(user *User) CreateUserUsecase {
	return &createUserUsecase{
		user: user,
	}
}

func (uc *createUserUsecase) Do() (string, error) {
	if uc.user.name != "" && uc.user.message != "" {
		return uc.user.message, nil
	} else {
		return "", errors.New("username or message is empty")
	}
}
