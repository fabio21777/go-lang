package main

import (
	"fmt"

	"github.com/fabio21777/bd/estrutura"
	"github.com/fabio21777/bd/insert"
	"github.com/fabio21777/bd/updateanddelete"
)

func main() {
	db := estrutura.AbrirConexaoComBancoDeDados()
	defer db.Close()
	estrutura.CriarDataBaseSeNaoExistir(db, "cursogo")
	estrutura.CriarTabelaUsuario(db)

	//inserts
	stmt := insert.CriarStmtUsuario(db)
	idJoao := insert.InsertUsuario(stmt, "João")
	idMaria := insert.InsertUsuario(stmt, "Maria")

	fmt.Println(idJoao, idMaria)

	defer stmt.Close()

	// transações
	//transacoes.Transacoes(db)

	// update e delete

	//update isso pode ser refatorado para não criar um novo statement é usar o mesmo do insert acima
	stmtUpdateAndDelete := updateanddelete.StmtUsuarioUpdateAndDelete(db, "UPDATE usuarios SET nome = $1 WHERE id = $2")

	updateanddelete.UpdateUsuario(stmtUpdateAndDelete, "João Silva", idJoao)
	updateanddelete.UpdateUsuario(stmtUpdateAndDelete, "Maria Silva", idMaria)

	stmtUpdateAndDelete = updateanddelete.StmtUsuarioUpdateAndDelete(db, "DELETE FROM usuarios WHERE id = $1")
	//deleta o joão
	updateanddelete.DeleteUsuario(stmtUpdateAndDelete, idJoao)

}
