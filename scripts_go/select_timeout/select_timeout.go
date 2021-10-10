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
			time.Sleep(time.Duration(rand.Intn(1100)) * time.Millisecond)
			c <- fmt.Sprintf("%v: %v", name, i)
		}
	}()

	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	c := Gen("Jane")

	timeout := time.After(5 * time.Second)

	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(time.Second):
			fmt.Println("You are too slow.")
			return
		case <-timeout:
			fmt.Println("5 seconds passed. I'm out.")
			return
		}
	}
}
