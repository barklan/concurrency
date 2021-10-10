package main

import (
	"fmt"
	"time"
)

func Gen(name string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(700 * time.Millisecond)
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
