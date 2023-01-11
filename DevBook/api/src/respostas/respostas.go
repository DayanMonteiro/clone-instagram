package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	primeiro parametro para escrever a resposta,

segundo o status da resposta
e o terceiro os dados mandados ao json em uma interface genérica
ou seja, JSON retorna uma resposta em JSON para a requisição
*/
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	// faz trazer a resposta como json no retorno do banco
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
