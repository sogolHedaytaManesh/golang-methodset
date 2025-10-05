package main

import "fmt"

type User struct {
}

func (u User) PrintA() {
	fmt.Println("PrintA with value receiver")
}

func (u *User) PrintB() {
	fmt.Println("PrintB with pointer receiver")
}
func main() {
	var u User
	u.PrintA()
	// u.PrintB() // ❌ Does not work
}
