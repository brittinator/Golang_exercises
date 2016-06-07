package main

import "fmt"

func main() {
    fmt.Print("Hello there! ")
    fmt.Println("Howdy")
    fmt.Print("The end")

    // initializing and declaring a var at the same time with :=
    f := "Brittz"
    fmt.Println(f)

    var lastname = "Walentin"
    fmt.Println(lastname)

    // loops: use for; that's it.

    i := 0
    for i < 10 {
      fmt.Println(i)
      i++
    }

    for j := 1; j < 10; j++ {
      fmt.Print("The number is: ")
      fmt.Println(j)
    }
}
