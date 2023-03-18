package contas

import (
	"github.com/Pedro-Cecilio/Projeto-Banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agência int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDeSaque float64) (result string) {
	if valorDeSaque <= c.saldo && valorDeSaque > 0 {
		c.saldo -= valorDeSaque
		return "Saque efetuado"
	} else if valorDeSaque > c.saldo {
		return "Saldo Insuficiente"
	} else {
		return "Valor de saque inválido"
	}
}
func (c *ContaCorrente) Depositar(ValorDeDeposito float64) (result string, Saldo float64) {
	if ValorDeDeposito > 0 {
		c.saldo += ValorDeDeposito
		return "Deposíto efetuado", c.saldo
	}
	return "Valor de depósito inválido", c.saldo
}
func (c *ContaCorrente) Transferir(contaAlvo *ContaCorrente, valorTransfe float64) (result string, saldoatt float64) {
	if c.saldo >= valorTransfe {
		if valorTransfe > 0 {
			c.saldo -= valorTransfe
			contaAlvo.Depositar(valorTransfe)
			return "Transferência realizada com sucesso.", c.saldo
		}
	} else {
		return "Saldo insuficiente", c.saldo
	}
	return "Valor de transferência inválido.", c.saldo
}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
