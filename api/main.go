package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Init()

	r := router.Create()

	portAPIStr := fmt.Sprintf(":%d", config.APIPort)

	fmt.Println("Rodando API na porta", portAPIStr)

	log.Fatal(http.ListenAndServe(portAPIStr, r))
}
