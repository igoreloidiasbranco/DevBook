package router

import "github.com/gorilla/mux"

func GerarRouter() *mux.Router {
	return mux.NewRouter()
}
