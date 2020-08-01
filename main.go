package main

import (
	"basic-go-kit-second/account"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db := GetDBconn()

	var (
		httpAddr = flag.String("http", ":8000", "http listen address")
	)
	rep := account.Repo{db}
	flag.Parse()
	ctx := context.Background()

	srv := account.AccountService{Repository: rep}
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoint := account.Endpoint{
		CreateCustomer:  account.MakeCreateCustomerEndpoint(srv),
		GetCustomerById: account.MakeGetCustomerByIdEndpoint(srv),
		GetAllCustomer:  account.MakeGetAllCustomerEndpoint(srv),
		UpdateCustomer:  account.MakeUpdateCustomerEndpoint(srv),
		DeleteCustomer:  account.MakeDeleteCustomerEndpoint(srv),
	}

	go func() {
		log.Println("basic go kit is listening on port:", *httpAddr)
		handler := account.NewHandler(ctx, endpoint)
		err := http.ListenAndServe(*httpAddr, handler)
		if err != nil {
			panic(err)
		}
	}()

	log.Fatalln(<-errChan)

}
