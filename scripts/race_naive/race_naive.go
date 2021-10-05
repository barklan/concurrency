package main

import (
	"fmt"
	"sync"
)

// multicore race -> one thread false safety
// -> breaking the safety with timer
func main() {
	// runtime.GOMAXPROCS(1)

	number := 0

	num_of_routines := 1_000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer wg.Done()

			local_number := number
			number = local_number + 1
		}()
	}

	wg.Wait()
	fmt.Println(number)
}
