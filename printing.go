package main

import "fmt"

func multiple(first, last string) (exclaim string) {
	fmt.Printf("Hello %v - %v", first, last)
	return "Hey"
}

func variableDeclared() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main() {
	m := multiple("B", "L")
	fmt.Print(m)

	fmt.Print("Declared variables but unused: ")
	variableDeclared()
}
