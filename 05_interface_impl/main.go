package main

import "fmt"

type Reader interface {
	Read()
}

type File struct{}

func (f *File) Read() {
	fmt.Println("Reading file...")
}

func main() {
	var f File
	// var r Reader = f // ❌ Error

	var r Reader = &f // ✅
	r.Read()
}
