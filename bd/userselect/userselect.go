package userselect

import (
	"database/sql"
	"log"
)

func SelectUsuario(db *sql.DB) {
	row, err := db.Query("SELECT id, nome FROM usuarios")
	if err != nil {
		log.Fatal("Erro ao selecionar usuários:", err)
	}
	defer row.Close()

	for row.Next() {
		var u usuario
		err := row.Scan(&u.id, &u.nome)
		if err != nil {
			log.Fatal("Erro ao escanear linha:", err)
		}
		log.Println("Usuário:", u)
	}

}

type usuario struct {
	id   int
	nome string
}
