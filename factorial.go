package main

import (
	"fmt"
)

func factorial(num int) int {
	// base case
	if num == 0 {
		fmt.Print("=")
		return 1
	}
	fmt.Printf("%v * ", num)
	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorial(25))
	// this doesn't work with numbers over 25, product becomes a negative number!
}
