// - Demonstrate with million goroutines
// -
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Println("took", time.Since(start))
	}
}

func main() {
	defer elapsed()()

	num_of_routines := 1_000

	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			fmt.Print("Hey!")
		}()
	}

	wg.Wait()
	fmt.Println("Done!")
}
