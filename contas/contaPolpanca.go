package contas

import (
	"github.com/Pedro-Cecilio/Projeto-Banco/clientes"
)
type ContaPolpanca struct {
	Titular clientes.Titular
	Conta, Agencia, Operacao int
	saldo float64

}

func (c *ContaPolpanca) Sacar(valorDeSaque float64) (result string) {
	if valorDeSaque <= c.saldo && valorDeSaque > 0 {
		c.saldo -= valorDeSaque
		return "Saque efetuado"
	} else if valorDeSaque > c.saldo {
		return "Saldo Insuficiente"
	} else {
		return "Valor de saque inválido"
	}
}
func (c *ContaPolpanca) Depositar(ValorDeDeposito float64) (result string, Saldo float64) {
	if ValorDeDeposito > 0 {
		c.saldo += ValorDeDeposito
		return "Deposíto efetuado", c.saldo
	}
	return "Valor de depósito inválido", c.saldo
}
func (c *ContaPolpanca) ObterSaldo() float64 {
	return c.saldo
}