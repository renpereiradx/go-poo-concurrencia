package main

func Sum(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// funcion de fibonacci mejorada
func Fibonacci2(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// otra funcion de fibonacci mejorada
// func ComputeOnceFibonacci(n int) int {
// 	size := math.Max(2, float64(n+1))
// 	values := make([]int, size)
// 	values[0] = 0
// 	values[1] = 1

// 	for i := 2; i <= n; i++ {
// 		values[i] = values[i-1] + values[i-2]
// 	}

// 	return values[n]
// }
// Code Coverage
// go test -v
// go test -v -cover
// go test -v -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go tool cover -func=coverage.out

// Profiling
// go test -bench=.
// go test -bench=. -benchmem
// go test -bench=. -benchmem -cpuprofile=cpu.out
// go tool pprof cpu.out
// go tool pprof cpu.out
// go tool pprof -http=:8080 cpu.out
// go test -bench=. -benchmem -memprofile=mem.out
// go tool pprof mem.out
// go tool pprof -http=:8080 mem.out
// go test -bench=. -benchmem -memprofile=mem.out -cpuprofile=cpu.out
