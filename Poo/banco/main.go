package main

import (
	"fmt"
	c "projetos/Poo/contas"
)

func main() {
	contaDaSilvia := c.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 600}

	status := contaDaSilvia.transferir(200, &contaDoGustavo)
	fmt.Println(status)
	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDoGustavo.Saldo)

	// fmt.Println(contaDaSilvia.Saldo)
	// fmt.Println(contaDaSilvia.sacar(200))
	// fmt.Println(contaDaSilvia.Saldo)
	// status, valor := contaDaSilvia.depositar(500)
	// fmt.Println(status, "\nSeu Saldo é:", valor)

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
