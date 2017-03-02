package main

import "fmt"

func telephone() func(word string) string {
	telephone := "I"
	return func(word string) string {
		telephone += word
		return telephone
	}
}

func main() {
	conversation := telephone()
	fmt.Println(conversation(" like"))
	fmt.Println(conversation(" to"))
	fmt.Println(conversation(" move"))
	fmt.Println(conversation(" it"))
	fmt.Println(conversation(" move"))
	fmt.Println(conversation(" it"))

}
