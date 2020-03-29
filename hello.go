package main

import (
	"bufio" //abrir arquivo
	"fmt"   //pacote para Println e scanf
	"io"
	"io/ioutil" //abrir arquivo
	"net/http"  //pacote para acessar internet
	"os"        //pacote para erros e sair do aplicativo
	"reflect"
	"strconv" //converter diversos tipos para string
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	// nome, idade := devolveNomeEIdade()
	// _, idade := devolveNomeEIdade()
	// fmt.Println(nome, idade)
	// exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
			imprimiLogs()
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
	//sites := []string{"https://www.alura.com.br", "https://www.google.com", "//https://www.gmail.com"}
	sites := leSitesDoArquivos()

	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}
	pulaLinha()

	//resp, err :=http.Get(site) //err caso aconteça algum erro

}

func testaSite(site string) {
	resp, err := http.Get(site) //err caso aconteça algum erro
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com erros. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func exibeNomes() {
	nomes := []string{"Caio", "Daniela", "Larissa", "Marcela"}
	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes))
}

func pulaLinha() {
	fmt.Println("")
}

func leSitesDoArquivos() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err :=ioutil.ReadFile("sites.txt")// bom para ler tudo de uma vez
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}
	leitor := bufio.NewReader(arquivo) //para ler linha por linha
	for {
		linha, err := leitor.ReadString('\n') //\n para ler ate o final
		linha = strings.TrimSpace(linha)      //para tirar os espacos do final do arquivo e tirar o /n

		sites = append(sites, linha)

		if err == io.EOF { //io.EOF erro de quando chega no final do arquivo/ ultima linha
			break //se encontrar o final do arquivo ele sai do for

		}

	}
	fmt.Println(sites)

	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //OpenFile consegue abrir arquivo, caso nao exista, ele pode ////abrir o arquivo e consegue escrever tambem, O_RDWR para pode escrever O_CREATE para poder criar caso nao tenha
	//O_APPEND para ele escrever na ultima linha

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "-online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimiLogs() {
	arquivo, err := ioutil.ReadFile("log.txt") //ele ja le o arquivo e ja fecha, entao nao precisa fechar
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
