package main

import (
	"fmt"
	"time"
)

func tarea(id int, texto string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("ID: %d - %s\n", id, texto)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go tarea(1, "Proceso A")
	go tarea(2, "Proceso B")
	go tarea(3, "Proceso C")

	time.Sleep(2 * time.Second)
}
