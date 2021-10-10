package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Msg struct {
	str  string
	wait chan bool // channels are first class objects
}

func Gen(name string) <-chan Msg {
	c := make(chan Msg)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- Msg{fmt.Sprintf("%v: %v", name, i), waitForIt}
			<-waitForIt // this blocks untils there is a value in there
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan Msg) <-chan Msg {
	c := make(chan Msg)

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
	fan := fanIn(Gen("Ann"), Gen("Joe"))

	for i := 0; i < 5; i++ {
		msg1 := <-fan
		fmt.Println(msg1.str)
		msg2 := <-fan
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}
}
