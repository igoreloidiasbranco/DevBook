package router

import (
	"api/router/rotas"

	"github.com/gorilla/mux"
)

func GerarRouter() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigurarRotas(r)
}
