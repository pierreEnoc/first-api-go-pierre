package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Livro struct {
	Id      int
	TItulo  string
	Autotor string
}

var Livros []Livro = []Livro{
	Livro{
		Id:      1,
		TItulo:  "O Guarani",
		Autotor: "Pierre",
	},

	Livro{
		Id:      2,
		TItulo:  "Camikase",
		Autotor: "Helena",
	},
	Livro{
		Id:      3,
		TItulo:  "Dom casmura",
		Autotor: "Machado de Assis",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
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
	http.ListenAndServe(":1337", nil) // DefaulServerMux
}

func main() {
	configurarSevidor()
}
