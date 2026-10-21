package main

import (
	"fmt"
	"sync"
)

func main() {
	contador := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(ptr *int) {
			defer wg.Done()
			*ptr++
		}(&contador)
	}

	wg.Wait()
	fmt.Println("Contador final:", contador)
}
