// Worker Pool
package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40, 41, 42, 43, 44}
	// tasks := []int{12, 40, 41, 42, 43, 44, 45, 46, 47, 48}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	// fmt.Println(len(jobs))
	close(jobs)
	// fmt.Println(len(jobs))

	for r := 0; r < len(tasks); r++ {
		fmt.Println(len(jobs))
		fmt.Println("Result: ", <-results)
		// <-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

/*
	Creo que la respuesta esta en el range de cada worker, este obtendra (cosumira) una tarea y despues pasara a ejecutar la serie fibonacci, quedandose en espera hasta obtener el resultado.

Es ahi donde entran los otros workers sin tareas a seguir consumiendo en el for range.

La manera en que se liberan, es en la ultima parte del main, cuando hacemos el ciclo for y solo vamos cosumiendo los resultados.


*/
