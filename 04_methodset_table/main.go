package main

type Demo struct {
}

func (d Demo) Foo() {

} // value receiver

func (d *Demo) Bar() {

} // pointer receiver

func main() {
	// var d Demo
	// d.Foo() ✅
	// d.Bar() ❌

	var dp *Demo
	dp.Foo() // ✅
	dp.Bar() // ✅
}
