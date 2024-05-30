package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Executando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
