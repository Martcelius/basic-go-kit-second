package main

import (
	"basic-go-kit-second/account"
	"basic-go-kit-second/router"
	"basic-go-kit-second/users"
	"context"
	"flag"
	"log"
	"net/http"
)

func main() {
	db := GetDBconn()
	router.InitRouter()
	defer db.Close()
	var (
		httpAddr = flag.String("http", ":8000", "http listen address")
	)

	repoAccount := account.Repo{Db: db}
	repoUser := users.UserRepo{Db: db}

	flag.Parse()
	ctx := context.Background()

	srvAccount := account.AccountService{Repository: repoAccount}
	srvUsers := users.UserService{RepoUser: repoUser}

	endpointAccount := account.Endpoint{
		CreateCustomer:  account.MakeCreateCustomerEndpoint(srvAccount),
		GetCustomerById: account.MakeGetCustomerByIdEndpoint(srvAccount),
		GetAllCustomer:  account.MakeGetAllCustomerEndpoint(srvAccount),
		UpdateCustomer:  account.MakeUpdateCustomerEndpoint(srvAccount),
		DeleteCustomer:  account.MakeDeleteCustomerEndpoint(srvAccount),
	}

	endpointUser := users.Endpoint{
		Register: users.MakeRegisterEndpoint(srvUsers),
	}

	log.Println("basic go kit is listening on port:", *httpAddr)
	account.NewHandlerAccount(ctx, endpointAccount)
	users.NewHandlerUser(ctx, endpointUser)
	err := http.ListenAndServe(*httpAddr, router.R)
	if err != nil {
		panic(err)
	}
}
