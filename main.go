package main

import (
	"fmt"
	"net/http"
)

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func configuraRotas() {
	http.HandleFunc("/", rotaPrincipal)

}

func configurarSevidor() {
	configuraRotas()

	fmt.Println("Servidor esta rodando na porta 1337")
	http.ListenAndServe(":1337", nil) // DefaulServerMux
}

func main() {
	configurarSevidor()
}
