package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	/*
	   O pacote ioutil do Go é uma biblioteca de utilitários de entrada/saída que fornece funções para trabalhar
	   com arquivos e diretórios. Algumas das funções mais comuns do ioutil incluem:

	   ReadFile: lê todo o conteúdo de um arquivo e o retorna como uma string ou um slice de bytes.
	   WriteFile: escreve o conteúdo de uma string ou slice de bytes em um arquivo.
	   ReadDir: lê o conteúdo de um diretório e retorna uma lista de entradas de arquivo.
	   TempDir: cria um novo diretório temporário no sistema de arquivos.
	   TempFile: cria um novo arquivo temporário no sistema de arquivos.
	   Essas funções são úteis para realizar tarefas básicas de entrada/saída,
	   como ler o conteúdo de um arquivo ou criar um novo arquivo temporário.
	*/

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	// variavel usuario é do tipo usuario
	var usuario usuario

	/*
	   Unmarshal é uma função do pacote "encoding/json" na linguagem de programação Go
	   que é usada para converter dados em formato JSON em uma estrutura de dados Go equivalente.
	   Ela é usada para ler dados em formato JSON de uma fonte (como uma string ou um arquivo)
	   e decodificá-los em um valor Go.
	*/
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		/*
		   "w.Write([]byte())" é uma função em Go que escreve os dados em forma de bytes em um "writer".
		   Um "writer" é um objeto que implementa a interface "io.Writer" e possui um método "Write"
		   que aceita um slice de bytes e retorna o número de bytes escritos e um erro (se houver).

		   A função "w.Write([]byte())" é comumente usada para escrever dados em uma conexão de rede,
		   um arquivo ou uma resposta HTTP
		*/
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

	fmt.Println(usuario)
}

// BuscarUsuarios traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// conectar com o banco
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	/*
	   O método Query de um objeto *sql.DB é usado para executar uma instrução SQL que retorna um conjunto de resultados.
	   Ele retorna um objeto *sql.Rows que pode ser usado para iterar sobre os resultados da consulta.
	*/
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	// executa uma iteração para cada linha retornada do banco com as informações dos usuários
	for linhas.Next() {
		var usuario usuario
		/*
		   O método Scan de um objeto *sql.Rows é usado para ler os valores de uma linha de um conjunto de resultados
		   retornado por um objeto *sql.Rows e armazená-los em variáveis.
		   Ele é usado em conjunto com o método Next, que é usado para avançar para a próxima linha do conjunto de resultados.
		*/
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	/*
	   Para respostas via http não usamos os métodos Marshall nem o Unmarshall nesse contexto é usado o método NewEncoder
	   O método NewEncoder da biblioteca encoding/json é usado para criar um novo objeto *Encoder que pode ser usado
	   para codificar valores em formato JSON.
	*/
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

// BuscarUsuario traz um usuário específico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}

}
