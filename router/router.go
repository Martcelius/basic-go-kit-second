package router

import (
	"github.com/gorilla/mux"
)

var R *mux.Router

func InitRouter() {
	R = mux.NewRouter()
}
