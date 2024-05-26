package estrutura

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func CriarDataBaseSeNaoExistir(db *sql.DB, dbName string) error {
	// Verifica se o banco de dados existe
	var name string
	query := "SELECT datname FROM pg_catalog.pg_database WHERE datname = $1"
	err := db.QueryRow(query, dbName).Scan(&name)
	if err == sql.ErrNoRows {
		// O banco de dados não existe, então cria
		_, err := db.Exec("CREATE DATABASE " + dbName) // Uso simples, mas cuidado com injeção de SQL
		if err != nil {
			return err
		}
		log.Println("Database created")
	} else if err != nil {
		return err
	} else {
		log.Println("Database already exists")
	}
	return nil
}

func CriarTabelaUsuario(db *sql.DB) {
	// Criação de tabela, deve  excluir se tabela usuario já existir
	Exec(db, "DROP TABLE IF EXISTS usuarios")
	Exec(db, `
	CREATE TABLE usuarios (
		id SERIAL PRIMARY KEY,
		nome TEXT
	);

`)
	log.Println("Tabela usuarios criada")
}

func AbrirConexaoComBancoDeDados() *sql.DB {
	db, err := sql.Open("postgres", "user=root password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
