package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 7)
	for i := 1; i <= 7; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(500 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println()
	fmt.Println("With bursty limiter:")

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 7)
	for i := 1; i <= 7; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}