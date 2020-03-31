package main

import (
	"net/http"

	"projetos/github.com/caio/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
