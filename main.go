package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

var Livros []Livro = []Livro{
	Livro{
		Id:     1,
		Titulo: "O Guarani",
		Autor:  "Pierre",
	},

	Livro{
		Id:     2,
		Titulo: "Camikase",
		Autor:  "Helena",
	},
	Livro{
		Id:     3,
		Titulo: "Dom casmura",
		Autor:  "Machado de Assis",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func configuraRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", listarLivros)

}

func configurarSevidor() {
	configuraRotas()

	fmt.Println("Servidor esta rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil)) // DefaulServerMux
}

func main() {
	configurarSevidor()
}
