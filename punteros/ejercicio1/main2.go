package main2

import "fmt"

func main() {
	numero := 42
	ptr := &numero

	fmt.Println("Valor:", numero)
	fmt.Println("Direccion:", ptr)
	fmt.Println("Valor del puntero:", *ptr)

	*ptr = 100

	fmt.Println("Nuevo valor:", numero)
}
