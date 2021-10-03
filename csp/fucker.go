package csp

import (
	"fmt"
)

func F1() {
	c := make(chan string)
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine says: %q\n", <-c)
	}

	fmt.Println("Exiting.")
}
