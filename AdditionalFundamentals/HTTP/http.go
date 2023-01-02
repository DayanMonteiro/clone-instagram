package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
	// Request - Response
	// Rotas
	// URI - Identificador do Recurso
	// Método - GET (busca dados), POST (cadastra dados), PUT (atualiza dados), DELETE (deleta dados)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // raiz
		w.Write([]byte("Página Raiz!"))
	})

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	// cria um servidor rodando na porta passada
	log.Fatal(http.ListenAndServe(":5000", nil))
}
