package main

import "fmt"

type Base struct {
}

func (b Base) PrintBase() {
	fmt.Println("Base method")
}

type Child struct {
	Base
}

func main() {
	c := Child{}
	c.PrintBase()
}
