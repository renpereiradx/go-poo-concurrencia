package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Started at#%d\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished at#%d\n", i)
	<-c
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)
	for i := 0; i < 5; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}
