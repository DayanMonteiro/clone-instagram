package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")

	r := router.Gerar()

	// subir o servidor
	log.Fatal(http.ListenAndServe(":5000", r))
}
