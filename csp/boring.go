package csp

import (
	"fmt"
	"time"
)

func boring(msg string, c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(1 * time.Second)
	}
}
