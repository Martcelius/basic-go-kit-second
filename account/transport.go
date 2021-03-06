package account

import (
	"basic-go-kit-second/router"
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateCustomerRequest
	err := json.NewDecoder(r.Body).Decode(&req.customer)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetCustomerById(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest

	id := mux.Vars(r)["id"]

	req = GetCustomerByIdRequest{Id: id}
	return req, nil
}

func decodeGetAllCustomer(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllCustomerRequest
	return req, nil
}

func decodeUpdateCustomer(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateCustomerRequest

	err := json.NewDecoder(r.Body).Decode(&req.customer)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteCustomer(_ context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]

	req := DeleteCustomerRequest{ID: id}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewHandlerAccount(ctx context.Context, endpoint Endpoint) {
	r := router.R

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoint.CreateCustomer,
		decodeCreateCustomerRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.GetCustomerById,
		decodeGetCustomerById,
		encodeResponse,
	))

	r.Methods("GET").Path("/user").Handler(httptransport.NewServer(
		endpoint.GetAllCustomer,
		decodeGetAllCustomer,
		encodeResponse,
	))

	r.Methods("PUT").Path("/user").Handler(httptransport.NewServer(
		endpoint.UpdateCustomer,
		decodeUpdateCustomer,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.DeleteCustomer,
		decodeDeleteCustomer,
		encodeResponse,
	))
}
