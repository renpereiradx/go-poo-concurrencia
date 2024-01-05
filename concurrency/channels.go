package main

/* Unbuffered channel:
Espera una función o una rutina para recibir el mensaje, es bloqueada por ser llamada en la misma función

Buffered channel:
Se puede llamar de manera inmediata, en el siguiente ejemplo 2 es el numero de canales que pueden ser usados */
import "fmt"

func printMessage(c chan string) {
	c <- "Hello World"
}

func main() {
	c := make(chan int, 2)
	c2 := make(chan string)
	go printMessage(c2)

	c <- 1
	fmt.Println(<-c)
	c <- 2
	fmt.Println(<-c)
	c <- 3
	fmt.Println(<-c)
	c <- 4
	fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c2)
}
