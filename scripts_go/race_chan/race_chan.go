package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(1)

	channel := make(chan int, 1)
	channel <- 0

	num_of_routines := 1_000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer wg.Done()

			local_number := <-channel
			time.Sleep(1 * time.Microsecond)
			channel <- local_number + 1
			// число = копия + 1
		}()
	}

	wg.Wait()
	fmt.Println(<-channel)
}
