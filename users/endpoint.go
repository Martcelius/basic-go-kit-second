package users

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type registerRequest struct {
	User users
}

type registerResponse struct {
	User interface{} `json:"user"`
	Msg  string      `json:"msg"`
	Err  string      `jsong:"err,omitempty"`
}

type loginRequest struct {
	User users
}

type loginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Err      string `json:"err"`
}

type Endpoint struct {
	Register endpoint.Endpoint
	Login    endpoint.Endpoint
}

func MakeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerRequest)
		userRegister, err := s.Register(ctx, req.User)

		if err != nil {
			return registerResponse{userRegister, "", err.Error()}, nil
		}
		return registerResponse{userRegister, "success register", ""}, nil
	}
}

func MakeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		fmt.Println("data endpoiint req", req)
		userLogin, err := s.Login(ctx, req.User)

		if err != nil {
			return loginResponse{"", "", "Email or Password Invalid"}, nil
		}

		return loginResponse{userLogin.Email, userLogin.Password, ""}, nil
	}
}
