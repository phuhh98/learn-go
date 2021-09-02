package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func (p *Person) speak() {

	fmt.Println("I'm", p.First, p.Last, ". I'm", p.Age)
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	p1 := Person{"James", "Bond", 34}

	saySomething(p1)

}
