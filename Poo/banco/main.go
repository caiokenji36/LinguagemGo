package main

import (
	"fmt"
	cont "projetos/Poo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := cont.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(50)
	PagarBoleto(&contaDoDenis, 30)

	contaDoCaio := cont.ContaCorrente{}
	contaDoCaio.Depositar(100)
	contaDoCaio.Sacar(50)
	PagarBoleto(&contaDoCaio, 10)

	contaDaPati := cont.ContaCorrente{}

	fmt.Println(contaDoDenis)
	fmt.Println(contaDoDenis.ObterSaldo())
	fmt.Println(contaDoCaio)
	fmt.Println(contaDaPati)

	//clienteBruno := cli.Titular{"Bruno", "123.31412.412", "Desenvolvedor"}
	//contaDoBruno := cont.ContaCorrente{clienteBruno, 123, 123141, 0}

	//fmt.Println(contaDoBruno)

}

func comentarios() {
	//contaDoCaio := ContaCorrente{titular: "Caio",
	//	numeroAgencia: 589, numeroConta: 123456, Saldo: 125.50}

	//contaDoCaio2 := ContaCorrente{titular: "Caio",
	//	numeroAgencia: 589, numeroConta: 123456, Saldo: 125.50}

	// contaDaLarissa := ContaCorrente{"Larissa",
	// 	5899, 133456, 145.50} // nao precisa passar os parametros

	//fmt.Println(contaDoCaio)
	//fmt.Println(contaDoCaio2 == contaDoCaio) // ele dar true, eles sao iguais
	// fmt.Println(contaDaLarissa)

	// var contaDaDani *ContaCorrente // outro tipo de criar um objeto, usando ponteiro
	// contaDaDani = new(ContaCorrente)
	// contaDaDani.titular = "Dani"
	// contaDaDani.Saldo = 500

	// fmt.Println(*contaDaDani)
}

func aulasAnteriores() {
	// contaDaSilvia := c.ContaCorrente{}
	// contaDaSilvia.Titular = "Silvia"
	// contaDaSilvia.Saldo = 500

	// contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 600}

	// status := contaDaSilvia.transferir(200, &contaDoGustavo)
	// fmt.Println(status)
	// fmt.Println(contaDaSilvia.Saldo)
	// fmt.Println(contaDoGustavo.Saldo)

	// fmt.Println(contaDaSilvia.Saldo)
	// fmt.Println(contaDaSilvia.sacar(200))
	// fmt.Println(contaDaSilvia.Saldo)
	// status, valor := contaDaSilvia.depositar(500)
	// fmt.Println(status, "\nSeu Saldo é:", valor)
}
