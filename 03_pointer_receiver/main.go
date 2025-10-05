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
	var up *User = &u

	up.PrintA() // ✅
	up.PrintB() // ✅
}
