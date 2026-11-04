package main

import (
	"fmt"
	"sync"
)

// Generador: envía números al canal de entrada
func generador(numeros ...int) <-chan int {
	entrada := make(chan int)

	go func() {
		defer close(entrada)

		for _, numero := range numeros {
			entrada <- numero
		}
	}()

	return entrada
}

// Fan-out: cada trabajador procesa números
func trabajador(entrada <-chan int, salida chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for numero := range entrada {
		cuadrado := numero * numero
		salida <- cuadrado
	}
}

func main() {
	// 1. Generar los números
	entrada := generador(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Canal centralizado de resultados
	resultados := make(chan int)

	// 2. Crear los trabajadores (fan-out)
	const cantidadTrabajadores = 3

	var wg sync.WaitGroup

	for i := 0; i < cantidadTrabajadores; i++ {
		wg.Add(1)
		go trabajador(entrada, resultados, &wg)
	}

	// 3. Fan-in: cerrar resultados cuando terminen
	// todos los trabajadores.
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// 4. Recopilar e imprimir los resultados
	fmt.Println("Cuadrados calculados:")

	for resultado := range resultados {
		fmt.Println(resultado)
	}
}
