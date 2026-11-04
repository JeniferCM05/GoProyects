package main

import (
	"fmt"
)

func main() {
	const nElements = 1_000_000
	arr := make([]int, nElements)
	for i := 0; i < nElements; i++ {

		arr[i] = 1
	}

	numGoroutines := 4
	chunkSize := nElements / numGoroutines

	results := make(chan int, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == numGoroutines-1 {
			end = nElements
		}

		go func(subSlice []int) {
			sumPartial := 0
			for _, val := range subSlice {
				sumPartial += val
			}
			results <- sumPartial
		}(arr[start:end])
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}

	fmt.Printf("La suma total del vector es: %d\n", totalSum)
}
