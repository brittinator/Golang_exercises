// this, with some typos and debugging, took me 8 minutes to complete on 12/19/2016

package main

import (
	"fmt"
	"strings"
)

func isPal(input string) bool {
	input = strings.ToUpper(input)
	end := len(input) / 2
	for i := 0; i <= end; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	answer := isPal("bob")
	fmt.Println(answer)
	answer = isPal("hey")
	fmt.Println(answer)
	answer = isPal("noon")
	fmt.Println(answer)
	answer = isPal("Noon")
	fmt.Println(answer)
}
