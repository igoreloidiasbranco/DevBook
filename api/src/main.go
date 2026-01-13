package main

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Api is running...")
	router := router.GerarRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
