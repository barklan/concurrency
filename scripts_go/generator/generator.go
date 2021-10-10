package main

import "fmt"

func generator(limit int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < limit; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func main() {
	for i := range generator(10) {
		fmt.Println(i)
	}
}
