package main

import "fmt"

type Account struct{}

func (a Account) Show() {
	fmt.Println("Show with value receiver")
}

func (a *Account) Update() {
	fmt.Println("Update with pointer receiver")
}

func main() {
	var acc Account
	acc.Show() // ✅
	// acc.Update() ❌

	var ap *Account = &acc
	ap.Show()   // ✅
	ap.Update() // ✅
}
