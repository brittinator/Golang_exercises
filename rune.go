package main

import "fmt"

func main() {
	myString := "Britt"
	for _, c := range myString {
		fmt.Println(string(c))
		if string(c) == "B" {
			fmt.Print("Equality")
		}
	}
}
