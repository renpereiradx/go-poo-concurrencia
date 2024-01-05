package main

import "fmt"

func Generator(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("For generator ", i)
		c <- i
	}
	close(c)
}

func Double(in chan int, out chan int) {
	for i := range in {
		fmt.Println("For double ", i)
		out <- i * 2
	}
	close(out)
}

func Print(c chan int) {
	for i := range c {
		println(i)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)
	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
