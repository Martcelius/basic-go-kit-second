package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	RepoErr             = errors.New("Unable to handle repo request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhoneNumNotFound = errors.New("Phone number not found")
)

// interface for interact repository
type repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomer(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateCustomer(ctx context.Context, customer Customer) error {
	_, err := r.Db.Exec("INSERT INTO customers (customerId, email, phone) VALUES (?,?,?)", customer.CustomerId, customer.Email, customer.Phone)

	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return err
	}

	return nil
}

func (r *Repo) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	var customer Customer
	err := r.Db.QueryRow("SELECT * FROM customers WHERE customerId=?", id).Scan(&customer.CustomerId, &customer.Email, &customer.Phone)

	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return customer, ErrIdNotFound
	}

	return customer, nil
}

func (r *Repo) GetAllCustomer(ctx context.Context) (interface{}, error) {
	var customer []Customer
	rows, err := r.Db.Query("SELECT customerId, email, phone FROM customers")

	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}

		return customer, err
	}

	defer rows.Close()
	for rows.Next() {
		each := Customer{}
		_ = rows.Scan(&each.CustomerId, &each.Email, &each.Phone)

		customer = append(customer, each)
	}

	return customer, nil
}

func (r *Repo) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	result, err := r.Db.Exec("UPDATE customers SET email=?, phone=? WHERE customerId=?", customer.Email, customer.Phone, customer.CustomerId)

	if err != nil {
		return "", nil
	}
	rowAffected, err := result.RowsAffected()

	if err != nil {
		return "", err
	}

	if rowAffected == 0 {
		return "", ErrIdNotFound
	}

	return "successfuly update", nil
}

func (r *Repo) DeleteCustomer(ctx context.Context, id string) (string, error) {
	result, err := r.Db.Exec("DELETE FROM customers WHERE customerId=?", id)

	if err != nil {
		return "", err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return "", err
	}

	if rowAffected == 0 {
		return "", ErrIdNotFound
	}

	return "successfuly delete", nil
}
