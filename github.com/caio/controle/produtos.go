package controle

var temp = template.Must(template.ParseGlob("templates/*.html")) //templates é a pasta onde esta os meus arquivos html

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) // esse Index é o que eu coloquei no começo da pagina index.html
}
