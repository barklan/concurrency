package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2, m3, m4 sync.Mutex
	перекресток := []*sync.Mutex{&m1, &m2, &m3, &m4}

	помеха := func(i int) int {
		if i == 3 {
			return 0
		} else {
			return i + 1
		}
	}

	машины := 4

	var wg sync.WaitGroup
	wg.Add(машины)

	for i := 0; i < машины; i++ {
		go func(дорога int) {
			defer wg.Done()

			перекресток[дорога].Lock()

			time.Sleep(time.Second)

			перекресток[помеха(дорога)].Lock()

			перекресток[дорога].Unlock()
			перекресток[помеха(дорога)].Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("All cars passed!")
}
