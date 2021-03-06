package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Gen(name string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("%v: %v", name, i)
		}
	}()

	return c
}

func main() {
	ann := Gen("Ann")
	joe := Gen("Joe")

	for i := 0; i < 5; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}

	fmt.Println("Bye!")
}
