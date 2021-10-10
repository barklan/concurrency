package main

import (
	"fmt"
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

	for i := 0; i < len(list); i++ {
		time.Sleep(time.Second)
		list[i].email = list[i].name + "@mail.com"
		fmt.Println(list)
	}
}
