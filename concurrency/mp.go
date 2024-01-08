package main

import (
	"time"
)

func doSomething(i time.Duration, c chan<- int, id int) {
	time.Sleep(i)
	c <- id
}

func main() {
	c := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c, 1)
	go doSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c:
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
		}
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c2)
}
