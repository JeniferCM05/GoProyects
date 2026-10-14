package main4

import "fmt"

type Libro struct {
	titulo string
	autor  string
	precio float64
}

func aplicarDescuento(libro *Libro, descuento float64) {
	libro.precio = libro.precio * (1 - descuento/100)
}

func main() {
	libro := Libro{
		titulo: "El principito",
		autor:  "Antoine de Saint-Exupery",
		precio: 200,
	}

	fmt.Println("Precio inicial:", libro.precio)

	aplicarDescuento(&libro, 20)

	fmt.Println("Precio final:", libro.precio)
}
