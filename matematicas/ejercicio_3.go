package main

import (
	"fmt"
	"sync"
)

type Job struct {
	id     int
	number int 
}

type Result struct {
	jobId  int
	result int64
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		var res int64 = 1
		for i := 1; i <= j.number; i++ {
			res *= int64(i)
		}
		results <- Result{jobId: j.id, result: res}
	}
}

func main() {
	const numJobs = 20
	const numWorkers = 4

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{id: j, number: j} 
	}
	close(jobs)s

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("Tarea ID: %d | Resultado Factorial: %d\n", res.jobId, res.result)
	}
}
