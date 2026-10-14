package main3

import "fmt"

func intercambiar(a *int, b *int) {

	temp := *a
	*a = *b
	*b = temp
}
func main() {

	a := 5
	b := 10

	fmt.Printf("Antes del intercambio: a = %d, b = %d\n", a, b)
	intercambiar(&a, &b)
	fmt.Printf("Después del intercambio: a = %d, b = %d\n", a, b)

}
