package main

import (
	"fmt"
)

func (s user) merge() string {
	return s.login + s.password
}

type user struct {
	login    string
	password string
}
type superUser struct {
	user
	hasPermit bool
}

type human interface {
	merge()
}

func main() {
	p1 := user{
		login:    "yuribenelli",
		password: "pippo",
	}
	p2 := user{
		login:    "giovannigalli",
		password: "pluto",
	}
	p3 := superUser{
		user: user{
			login:    "mariorossi",
			password: "paperino",
		},
		hasPermit: true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println("PERSON 1", p1.merge())
	fmt.Println("PERSON 2", p2.merge())
	fmt.Println("PERSON 3", p3.merge())

	fmt.Printf("PERSON 1: \t%v\t%T\n", p1, p1)
	fmt.Printf("PERSON 1: \t%v\t%T\n", p2, p2)
	fmt.Printf("PERSON 1: \t%v\t%T\n", p3, p3)
}
