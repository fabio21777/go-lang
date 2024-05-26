package insert

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertUsuario(stmt *sql.Stmt, name string) int {
	var lastInsertId int
	// Usar QueryRow para executar a inserção e capturar o ID retornado
	err := stmt.QueryRow(name).Scan(&lastInsertId)
	if err != nil {
		log.Fatal("Erro ao inserir usuário e obter ID:", err)
	}
	fmt.Println("Usuário inserido -", name, "com ID:", lastInsertId)
	return lastInsertId
}

func CriarStmtUsuario(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("INSERT INTO usuarios (nome) VALUES ($1) RETURNING id")
	if err != nil {
		log.Fatal("Erro ao preparar statement:", err)
	}
	return stmt
}
