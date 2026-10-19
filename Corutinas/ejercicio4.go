package main

import (
	"fmt"
)

func sumar(a int, b int, ch chan int) {
	ch <- a + b
}

func main() {
	ch := make(chan int)
	go sumar(15, 25, ch)
	fmt.Println(<-ch)
}
