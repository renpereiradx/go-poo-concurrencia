package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started at#%d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished at#%d\n", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
}
