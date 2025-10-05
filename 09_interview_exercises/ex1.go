package main

import "fmt"

type Sample struct{}

func (s Sample) A() {
	fmt.Println("A")
}

func (s *Sample) B() {
	fmt.Println("B")
}

func main() {
	var s Sample
	var sp *Sample = &s

	s.A() // ✅
	// s.B() ❌
	sp.A() // ✅
	sp.B() // ✅
}
