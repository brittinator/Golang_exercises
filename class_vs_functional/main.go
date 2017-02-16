package main

import "fmt"

type A string

func (a A) Demo(i int) string {
	fmt.Println("string")
	return fmt.Sprintf("%s-%d", a, i)
}
func Demo(a A, i int) string {
	fmt.Println("string")
	return fmt.Sprintf("%s-%d", a, i)
}
