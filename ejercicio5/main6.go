package main6

import "fmt"

type CuentaBancaria struct {
	saldo float64
}

func (c *CuentaBancaria) Depositar(monto float64) {
	if monto > 0 {
		c.saldo += monto
	}
}

func (c *CuentaBancaria) Retirar(monto float64) {
	if monto <= 0 {
		fmt.Println("Error: monto no valido")
		return
	}

	if monto > c.saldo {
		fmt.Println("Error: saldo insuficiente")
		return
	}

	c.saldo -= monto
}

func main() {
	cuenta := CuentaBancaria{saldo: 4000}

	cuenta.Depositar(400)
	cuenta.Retirar(200)

	fmt.Println("Saldo final:", cuenta.saldo)
}
