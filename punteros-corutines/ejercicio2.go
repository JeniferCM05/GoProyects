package main

import (
	"fmt"
	"sync"
)

type Resultado struct {
	Suma       int
	Completado bool
}

func procesar(res *Resultado, numeros []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, n := range numeros {
		res.Suma += n
	}

	res.Completado = true
}

func main() {
	res := Resultado{}
	numeros := []int{10, 20, 30, 40, 50}

	var wg sync.WaitGroup
	wg.Add(1)

	go procesar(&res, numeros, &wg)

	wg.Wait()

	fmt.Println("Suma:", res.Suma)
	fmt.Println("Completado:", res.Completado)
}
