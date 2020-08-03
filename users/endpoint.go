package users

import (
	"context"

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

type Endpoint struct {
	Register endpoint.Endpoint
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
