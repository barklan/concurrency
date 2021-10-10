package main

import (
	"math/rand"
	"sync"
)

func main() {
	n := 1_000_000

	a := make([]int, n)

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			a[j] = rand.Intn(100) * rand.Intn(100)
		}(i)
	}

	wg.Wait()

	// fmt.Println(a)
}
