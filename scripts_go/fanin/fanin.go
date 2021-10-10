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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	ann := Gen("Ann")
	joe := Gen("Joe")

	fan := fanIn(ann, joe)

	for i := 0; i < 10; i++ {
		fmt.Println(<-fan)
	}

	fmt.Println("Bye!")
}
