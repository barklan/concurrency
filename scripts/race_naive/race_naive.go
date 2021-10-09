package main

import (
	"fmt"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(1) // these are os threads

	number := 0

	num_of_routines := 1000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer func() {
				wg.Done()
				m.Unlock()
			}()
			m.Lock()
			local_number := number
			number = local_number + 1
		}()
	}

	wg.Wait()
	fmt.Println(number)
}
