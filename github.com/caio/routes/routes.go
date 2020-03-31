package routes

import (
	"net/http"

	"projetos/github.com/caio/controle"
)

func CarregaRotas() {
	http.HandleFunc("/", controle.Index)
	http.HandleFunc("/new", controle.New)
	http.HandleFunc("/insert", controle.Insert)
	http.HandleFunc("/delete", controle.Delete)
	http.HandleFunc("/edit", controle.Edit)
	http.HandleFunc("/update", controle.Update)
}
