package main

import (
	
	"fmt"

	"github.com/Pedro-Cecilio/Projeto-Banco/clientes"
	"github.com/Pedro-Cecilio/Projeto-Banco/contas"
)

func main() {
	cliente1 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Pedro",
			Cpf:       "14815587655",
			Profissao: "Desenvolvedor Go",
		},
		Agência: 5566,
		Conta:   145682,
	}
	
	fmt.Println(cliente1)
}
