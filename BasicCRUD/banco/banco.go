package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "nome:senha@/nome-do-banco?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	/*
	   O comando Ping() na linguagem Go é uma função que faz parte do pacote net,
	   e é utilizado para verificar se um determinado endereço de rede pode ser alcançado.
	   Quando você chama a função ping, ela envia um pacote de dados para o endereço especificado e aguarda uma resposta.
	   Se a resposta for recebida, isso indica que o endereço de rede pode ser alcançado e que a conexão é funcional.
	   Se não for recebida, isso pode indicar um problema com a conexão de rede.
	*/

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
