package updateanddelete

import (
	"database/sql"
	"log"
)

func UpdateUsuario(stmt *sql.Stmt, nome string, id int) {
	stmt.Exec(nome, id)
}

func DeleteUsuario(stmt *sql.Stmt, id int) {
	stmt.Exec(id)
}

func StmtUsuarioUpdateAndDelete(db *sql.DB, prepare string) *sql.Stmt {
	stmt, err := db.Prepare(prepare)
	if err != nil {
		log.Fatal("Erro ao preparar statement:", err)
	}
	return stmt
}
