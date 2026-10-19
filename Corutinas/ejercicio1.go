package maim

import (
	"fmt"
	"time"
)

func miFuncion() {
	fmt.Println("¡Hola desde una goroutine!")
}

func main() {
	go miFuncion()
	time.Sleep(1 * time.Second)
}
