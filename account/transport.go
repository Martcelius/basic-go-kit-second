package account

import (
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

	req := DeleteCustomerRequest{Id: id}
	return req, nil
}

func encodeReponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func NewHandler(ctx context.Context, endpoint Endpoint) http.Handler {
	r := mux.NewRouter()

	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoint.CreateCustomer,
		decodeCreateCustomerRequest,
		encodeReponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.GetCustomerById,
		decodeGetCustomerById,
		encodeReponse,
	))

	r.Methods("GET").Path("/user").Handler(httptransport.NewServer(
		endpoint.GetAllCustomer,
		decodeGetAllCustomer,
		encodeReponse,
	))

	r.Methods("PUT").Path("/user").Handler(httptransport.NewServer(
		endpoint.UpdateCustomer,
		decodeUpdateCustomer,
		encodeReponse,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.DeleteCustomer,
		decodeDeleteCustomer,
		encodeReponse,
	))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
