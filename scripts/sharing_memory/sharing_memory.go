package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			a++
		}()
	}

	wg.Wait()

	fmt.Print(a)
}
