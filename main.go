package main

import (
	"fmt"

	"github.com/Pedro-Cecilio/Projeto-Banco/clientes"
	"github.com/Pedro-Cecilio/Projeto-Banco/contas"
)
// Interface que implementa structs que possuam o método Sacar()
type VerificarConta interface {
	Sacar(valorDeSaque float64) string
}
// Função para pagamento de boleto. Qualquer conta que for implementada pela interface VerificarConta pode pagar boleto.
func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}
func main() {
	cliente1 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Pedro",
			Cpf:       "14815587655",
			Profissao: "Desenvolvedor",
		},
		Agência: 5566,
		Conta:   145682,
	}
	cliente2 := contas.ContaPolpanca{
		Titular: clientes.Titular{
			Nome:      "Marta",
			Cpf:       "83728383022",
			Profissao: "Desenvolvedor",
		},
		Agencia: 4432,
		Conta:   890123,
	}
	cliente3 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Natan",
			Cpf:       "14814745625",
			Profissao: "Desenvolvedor",
		},
		Agência: 5566,
		Conta:   152682,
	}
	cliente1.Depositar(5559.95)
	cliente2.Depositar(5559.95)
	cliente3.Depositar(5559.95)

	PagarBoleto(&cliente1, 59.90)
	PagarBoleto(&cliente2, 655.33)
	PagarBoleto(&cliente3, 855.33)

	fmt.Println(cliente1.ObterSaldo())
	fmt.Println(cliente2.ObterSaldo())
	fmt.Println(cliente3.ObterSaldo())

	cliente1.Transferir(&cliente3, 427.66)
	fmt.Println(cliente1.ObterSaldo())
	fmt.Println(cliente3.ObterSaldo())
}
