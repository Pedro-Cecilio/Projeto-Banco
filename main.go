package main

import (
	// "Banco/clientes"
	"Banco/contas"
	"fmt"
)

func main() {
	cliente1 := contas.ContaCorrente{
		// Titular: clientes.Titular{
		// 	Nome:      "Pedro",
		// 	Cpf:       "14815587655",
		// 	Profissao: "Desenvolvedor Go",
		// },
		Agência: 5566,
		Conta:   145682,
		Saldo:   125550.55,
	}

	fmt.Println(cliente1)
}
