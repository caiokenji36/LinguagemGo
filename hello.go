package main

import (
	"fmt" //pacote para Println e scanf
	"net/http"

	//pacote para acessar internet
	"os" //pacote para erros e sair do aplicativo
	"reflect"
)

func main() {
	nome, idade := devolveNomeEIdade()
	//_, idade := devolveNomeEIdade()
	fmt.Println(nome, idade)
	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco essa")
			os.Exit(-1)
		}
	}

	//if comando == 1 {
	//	fmt.Println("iniciando monitoramento")
	//} else if comando == 2 {
	//	fmt.Println("Exibindo logs")
	//} else if comando == 0 {
	//	fmt.Println("Saindo do programa")
	//} else {
	//	fmt.Println("Nao conheco essa")
	//}

}

func devolveNomeEIdade() (string, int) {
	nome := "Caio"
	idade := 24
	return nome, idade
}

func exibeIntroducao() {
	var nome string = "Caio"
	var dinheiro float32 = 23.5
	var idade int = 23

	fmt.Println("Ola mundo, caio")
	fmt.Println("Meu primeiro programa em Go")
	fmt.Println("Bem vindo ", nome)
	fmt.Println("Ola", nome, "sua idade é de ", idade, "seu saldo é de", dinheiro)
	fmt.Println("Bem vindo ", reflect.TypeOf(nome))
}

func exibeMenu() {
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Sair do Programa")
}

func leComando() int { //o retorno da funcao fica depois
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("iniciando monitoramento")
	site := "https://www.alura.com.br"
	//resp, err :=http.Get(site) //err caso aconteça algum erro
	resp, _ := http.Get(site) //err caso aconteça algum erro

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "está com erros. Status Code:", resp.StatusCode)
	}

}
