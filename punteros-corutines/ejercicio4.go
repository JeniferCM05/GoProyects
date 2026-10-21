package main

import (
	"fmt"
)

type Tarea struct {
	ID     int
	Precio float64
}

func worker(tareas chan *Tarea) {
	for t := range tareas {
		t.Precio = t.Precio * 1.16
	}
}

func main() {
	tareas := make(chan *Tarea, 5)

	t1 := Tarea{ID: 1, Precio: 100}
	t2 := Tarea{ID: 2, Precio: 200}

	go worker(tareas)

	tareas <- &t1
	tareas <- &t2

	close(tareas)

	fin := make(chan bool)

	go func() {
		worker(tareas)
		fin <- true
	}()

	<-fin

	fmt.Println("Tarea 1:", t1.ID, t1.Precio)
	fmt.Println("Tarea 2:", t2.ID, t2.Precio)
}
