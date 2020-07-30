package account

import (
	"context"
	"fmt"
)

type Customer struct {
	CustomerId string `json: "customerId"`
	Email      string `json: "email"`
	Phone      string `json: "email"`
}

type Service interface {
	CreateCustomer(ctx context.Context, customer Customer) (string, error)
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomer(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}

type AccountService struct {
	repository Repo
}

//Implementation interface AccountService

func (s AccountService) CreateCustomer(ctx context.Context, customer Customer) (string, error) {
	customerDetail := Customer{
		CustomerId: customer.CustomerId,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	err := s.repository.CreateCustomer(ctx, customerDetail)
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	return "Success", nil
}

func (s AccountService) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	customer, err := s.repository.GetCustomerById(ctx, id)

	if err != nil {
		fmt.Println("error", err)
		var empty interface{}
		return empty, err
	}

	return customer, nil
}

func (s AccountService) GetAllCustomer(ctx context.Context) (interface{}, error) {
	customer, err := s.repository.GetAllCustomer(ctx)

	if err != nil {
		fmt.Println("error", err)
		var empty interface{}
		return empty, err
	}

	return customer, nil
}

func (s AccountService) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	customerUpdate := Customer{
		CustomerId: customer.CustomerId,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	msg, err := s.repository.UpdateCustomer(ctx, customerUpdate)
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	return msg, nil
}

func (s AccountService) DeleteCustomer(ctx context.Context, id string) (string, error) {
	msg, err := s.repository.DeleteCustomer(ctx, id)

	if err != nil {
		fmt.Println("error", err)
		return msg, err
	}

	return msg, nil
}
