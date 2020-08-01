package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// struct for mapping incoming request and response

type CreateCustomerRequest struct {
	customer Customer
}

type CreateCustomerResponse struct {
	Msg string `json:"msg"`
	Err error  `jsong:"err,omitempty"`
}

type GetCustomerByIdRequest struct {
	Id string `json:"customerId"`
}

type GetCustomerByIdResponse struct {
	Customer interface{} `json:"customer,omitempty"`
	Err      string      `json:"err,omitempty"`
}

type GetAllCustomerRequest struct{}

type GetAllCustomerResponse struct {
	Customer interface{} `json:"customer,omitempty"`
	Err      string      `json:"err,omitempty"`
}

type DeleteCustomerRequest struct {
	ID string `json:"customerId"`
}

type DeleteCustomerResponse struct {
	Msg string `json:"msg"`
	Err error  `json:"err,omitempty"`
}

type UpdateCustomerRequest struct {
	customer Customer
}

type UpdateCustomerResponse struct {
	Msg string `json:"status,omitempty"`
	Err error  `json:"error,omitempty"`
}

// implement our endpoint method to consume request and response struct

type Endpoint struct {
	CreateCustomer  endpoint.Endpoint
	GetCustomerById endpoint.Endpoint
	GetAllCustomer  endpoint.Endpoint
	UpdateCustomer  endpoint.Endpoint
	DeleteCustomer  endpoint.Endpoint
}

func MakeCreateCustomerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		msg, err := s.CreateCustomer(ctx, req.customer)

		if err != nil {
			return CreateCustomerResponse{"", err}, nil
		}

		return CreateCustomerResponse{Msg: msg, Err: err}, nil
	}
}

func MakeGetCustomerByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIdRequest)

		customer, err := s.GetCustomerById(ctx, req.Id)

		if err != nil {
			return GetCustomerByIdResponse{customer, "Error Get Customer ByI Id endpoint"}, nil
		}
		return GetCustomerByIdResponse{customer, ""}, nil
	}
}

func MakeGetAllCustomerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customerDetail, err := s.GetAllCustomer(ctx)

		if err != nil {
			return GetAllCustomerResponse{Customer: customerDetail, Err: "Data Not Found"}, nil
		}
		return GetAllCustomerResponse{Customer: customerDetail, Err: ""}, nil
	}
}

func MakeUpdateCustomerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCustomerRequest)
		msg, err := s.UpdateCustomer(ctx, req.customer)

		if err != nil {
			return UpdateCustomerResponse{"Error update", err}, nil
		}

		return UpdateCustomerResponse{msg, err}, nil
	}
}

func MakeDeleteCustomerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCustomerRequest)
		msg, err := s.DeleteCustomer(ctx, req.ID)

		if err != nil {
			return DeleteCustomerResponse{"Error update", err}, nil
		}

		return DeleteCustomerResponse{msg, err}, nil
	}
}
