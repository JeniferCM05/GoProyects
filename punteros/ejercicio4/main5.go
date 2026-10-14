package main5

import "fmt"

func duplicar(ptr *int) {
	if ptr == nil {
		fmt.Println("Error: puntero nulo")
		return
	}

	*ptr *= 2
	fmt.Println("Valor duplicado:", *ptr)
}

func main() {
	numero := 10

	duplicar(&numero)
	fmt.Println("Valor de numero:", numero)

	duplicar(nil)
}
