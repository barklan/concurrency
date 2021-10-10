package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	name  string
	email string
}

func main() {
	list := []User{
		{"John", ""},
		{"Jim", ""},
		{"Jane", ""},
		{"Glebushek", ""},
	}

	var wg sync.WaitGroup
	wg.Add(len(list))

	for i := 0; i < len(list); i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			list[i].email = list[i].name + "@mail.com"
			fmt.Println(list)
		}(i)
	}

	wg.Wait()
}
