package main

import "fmt"

type Person struct{}

func (p Person) SayHello() {
	fmt.Println("Hi from value receiver")
}
func (p *Person) Walk() {
	fmt.Println("Walking from pointer receiver")
}

func main() {
	var p Person
	p.SayHello()
	// p.Walk() // ❌ Error

	var pp = &p
	pp.SayHello()
	pp.Walk()
}
