package client

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

/*
	é necessário que a primeira letra seja maiúscula para que o json.Marshal consiga acessar os campos, caso ao contrário, ele não consegue acessar os campos
*/
type usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	log.Println("executei")
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == http.MethodGet && id > 0:
		usuarioPorId(w, r, id)
	case r.Method == http.MethodGet:
		usuariosTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Deu ruim")
	}

}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("postgres", "user=root password=root sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = $1", id).Scan(&usuario.ID, &usuario.Nome)
	json, _ := json.Marshal(usuario)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(json))
	fmt.Fprint(w, string(json))
}

func usuariosTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=root password=root sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		//fmt.Println(usuario)
		usuarios = append(usuarios, usuario)
	}
	fmt.Println(usuarios)
	json, err := json.Marshal(usuarios)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--->", string(json))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}
