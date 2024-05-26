package transacoes

import (
	"database/sql"
	"log"
)

func Transacoes(db *sql.DB) {

	// iniciar transação
	tx, _ := db.Begin()
	// criar statement
	stmt, _ := tx.Prepare("INSERT INTO usuarios (id, nome) VALUES ($1, $2)")

	// executar statement
	stmt.Exec(2000, "João")
	stmt.Exec(2001, "Maria")

	// executar statement com erro
	_, err := stmt.Exec(1, "Pedro")

	// verificar erro, se houver, desfazer transação
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
