package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "Message."
	c <- "Second message."

	close(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
}
