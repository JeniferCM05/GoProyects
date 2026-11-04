package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const size = 200

	A := make([][]int, size)
	B := make([][]int, size)
	CSeq := make([][]int, size)
	CConc := make([][]int, size)

	for i := 0; i < size; i++ {
		A[i] = make([]int, size)
		B[i] = make([]int, size)
		CSeq[i] = make([]int, size)
		CConc[i] = make([]int, size)
		for j := 0; j < size; j++ {
			A[i][j] = i + j
			B[i][j] = i - j
			CSeq[i][j] = 0
			CConc[i][j] = 0
		}
	}

	startSeq := time.Now()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				CSeq[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	durationSeq := time.Since(startSeq)
	fmt.Printf("Tiempo Secuencial: %s\n", durationSeq)

	startConc := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					CConc[row][j] += A[row][k] * B[k][j]
				}
			}
		}(i)
	}

	wg.Wait()
	durationConc := time.Since(startConc)
	fmt.Printf("Tiempo Concurrente: %s\n", durationConc)
}
