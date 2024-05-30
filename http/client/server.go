package main

import (
	"log"
	"net/http"

	"github.com/fabio21777/server/client"
)

func main() {
	http.HandleFunc("/usuarios/", client.UsuarioHandler)
	log.Println("Executando na porta 3000...")
	http.ListenAndServe(":3000", nil)
}
