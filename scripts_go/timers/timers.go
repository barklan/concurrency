package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(3 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired.")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired.")
	}()

	timer2.Stop()
	fmt.Println("Timer 2 stopped.")

	time.Sleep(time.Second * 2)
}
