package main

import (
	"api/config"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarVariaveisAmbiente()

	router := router.GerarRouter()

	fmt.Printf("Api is running on port %d\n", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
