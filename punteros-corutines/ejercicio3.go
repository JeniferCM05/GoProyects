package main

import (
	"fmt"
	"sync"
)

func main() {
	datos := make([]int, 5)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			p := &datos[i]
			*p = i * 10
		}(i)
	}

	wg.Wait()

	fmt.Println("Datos:", datos)
}
