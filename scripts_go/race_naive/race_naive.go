package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(4) // these are os threads

	number := 0

	num_of_routines := 1000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()

			b := number
			time.Sleep(time.Second)
			b++
			number = b
		}()
	}

	wg.Wait()
	fmt.Println(number)
}
