package users

import (
	"basic-go-kit-second/router"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req registerRequest
	fmt.Println("<------------decode")
	err := json.NewDecoder(r.Body).Decode(&req.User)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req.User)

	if err != nil {
		return nil, err
	}
	fmt.Printf("data decode", req)
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("<------------encode")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewHandlerUser(ctx context.Context, endpoint Endpoint) {

	r := router.R
	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoint.Register,
		decodeRegisterRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/login").Handler(httptransport.NewServer(
		endpoint.Login,
		decodeLoginRequest,
		encodeResponse,
	))
}
