package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	router := mux.NewRouter()
	return rotas.Configurar(router)
}
