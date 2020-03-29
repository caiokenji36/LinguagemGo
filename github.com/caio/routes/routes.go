package routes

import (
	"github.com/caio/controle"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controle.Index)
}
