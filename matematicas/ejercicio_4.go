package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	numGoroutines = 10
	incrementos   = 100000
)

// Variables globales compartidas
var contador int
var contadorMutex int
var contadorAtomic int64

var mutex sync.Mutex

// 1. Versión SIN protección
func incrementarSinProteccion(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < incrementos; i++ {
		contador++
	}
}

// 2. Versión con sync.Mutex
func incrementarConMutex(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < incrementos; i++ {
		mutex.Lock()
		contadorMutex++
		mutex.Unlock()
	}
}

// 3. Versión con sync/atomic
func incrementarAtomicamente(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < incrementos; i++ {
		atomic.AddInt64(&contadorAtomic, 1)
	}
}

func main() {
	var wg sync.WaitGroup

	// PRUEBA 1: Sin protección
	fmt.Println("=== SIN PROTECCION ===")

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarSinProteccion(&wg)
	}

	wg.Wait()

	fmt.Println("Resultado:", contador)
	fmt.Println("Esperado: ", numGoroutines*incrementos)

	// PRUEBA 2: Con Mutex
	fmt.Println("\n=== CON SYNC.MUTEX ===")

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarConMutex(&wg)
	}

	wg.Wait()

	fmt.Println("Resultado:", contadorMutex)
	fmt.Println("Esperado: ", numGoroutines*incrementos)

	// PRUEBA 3: Con Atomic
	fmt.Println("\n=== CON SYNC/ATOMIC ===")

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarAtomicamente(&wg)
	}

	wg.Wait()

	fmt.Println("Resultado:", contadorAtomic)
	fmt.Println("Esperado: ", numGoroutines*incrementos)
}
