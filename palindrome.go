package main

import "fmt"

func isPalindrome(input string) bool {
	fmt.Println(input)
	beginning := input[0]
	end := input[len(input)-1]
	fmt.Println("%v, %v", beginning, end)
	if len(input) <= 1 {
		return true
	}

	if beginning != end {
		return false
	}
	isPalindrome(input[beginning+1 : end-1])
	return true
}

func main() {
	answer := isPalindrome("racecar")
	fmt.Println(answer)
}
