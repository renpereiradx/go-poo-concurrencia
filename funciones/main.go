package main

import (
	"fmt"
	"time"
)

func ObtenerValores(numero int) (doble int, triple int, cuadruple int) {
	doble = numero * 2
	triple = numero * 3
	cuadruple = numero * 4
	return
}

// funcion variatica
func Suma(numeros ...int) (total int) {
	for _, valor := range numeros {
		total += valor
	}
	return
}

// funcion variatica
func ImprimirNombres(nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

func main() {
	function := func() string {
		return "Hello world"
	}
	fmt.Println(function())
	function2 := func(name string) string {
		return fmt.Sprintf("Hello %s", name)
	}
	fmt.Println(function2("John"))
	function3 := func() string {
		return fmt.Sprintf("Hello %s", "John")
	}()
	fmt.Println(function3)
	function4 := func(name string) string {
		return fmt.Sprintf("Hello %s", name)
	}("Carlos")
	fmt.Println(function4)
	func() {
		fmt.Println("Hello world anonymous function")
	}()

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(2 * time.Second)
		fmt.Println("Finishing function")
		c <- 1
	}()
	fmt.Println(<-c)
	fmt.Println(Suma(1, 2, 3, 4))
	ImprimirNombres("John", "Jane", "Doe")
	doble, triple, cuadruple := ObtenerValores(2)
	fmt.Println(doble, triple, cuadruple)
}
