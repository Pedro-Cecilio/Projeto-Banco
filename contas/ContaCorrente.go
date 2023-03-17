package contas

import (
	"Banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agência int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(valorDeSaque float64) (result string) {
	if valorDeSaque <= c.Saldo && valorDeSaque > 0 {
		c.Saldo -= valorDeSaque
		return "Saque efetuado"
	} else if valorDeSaque > c.Saldo {
		return "Saldo Insuficiente"
	} else {
		return "Valor de saque inválido"
	}
}
func (c *ContaCorrente) Depositar(ValorDeDeposito float64) (result string, Saldo float64) {
	if ValorDeDeposito > 0 {
		c.Saldo += ValorDeDeposito
		return "Deposíto efetuado", c.Saldo
	}
	return "Valor de depósito inválido", c.Saldo
}
func (c *ContaCorrente) Transferir(contaAlvo *ContaCorrente, valorTransfe float64) (result string, saldoatt float64) {
	if c.Saldo >= valorTransfe {
		if valorTransfe > 0 {
			c.Saldo -= valorTransfe
			contaAlvo.Depositar(valorTransfe)
			return "Transferência realizada com sucesso.", c.Saldo
		}
	} else {
		return "Saldo insuficiente", c.Saldo
	}
	return "Valor de transferência inválido.", c.Saldo
}
