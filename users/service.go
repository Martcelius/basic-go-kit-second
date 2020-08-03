package users

import (
	"context"

	validationEmail "github.com/mcnijman/go-emailaddress"
)

type users struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type Service interface {
	Register(ctx context.Context, u users) (interface{}, error)
	Login(ctx context.Context, u users) (users, error)
}

type UserService struct {
	RepoUser UserRepo
}

// implement interface

func (user UserService) Register(ctx context.Context, u users) (interface{}, error) {
	_, err := validationEmail.Parse(u.Email)
	if err != nil {
		return user, err
	}

	userRegister := users{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Password:  u.Password,
	}

	err = user.RepoUser.Register(ctx, userRegister)

	if err != nil {
		return users{}, err
	}
	return userRegister, nil
}

func (user UserService) Login(ctx context.Context, u users) (users, error) {
	userLogin := users{
		Email:    u.Email,
		Password: u.Password,
	}

	_, err := validationEmail.Parse(userLogin.Email)

	if err != nil {
		return u, err
	}

	result, err := user.RepoUser.Login(ctx, userLogin)

	if err != nil {
		return result, err
	}

	return result, nil
}
