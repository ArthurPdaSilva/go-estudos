package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// html dinâmico
var templates *template.Template

type User struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	//Criando uma rota
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := User{"João", "pedro@gmail.com"}
		templates.ExecuteTemplate(w, "index.html", u)
		w.Write([]byte("Olá mundo!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página de usuários!"))
	})

	//Criando servidor na porta 5000
	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
