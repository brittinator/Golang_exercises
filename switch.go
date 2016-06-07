package main

import "fmt"
import "time"

func main() {
  age :=31
  switch true {
  case age < 18:
    fmt.Println("You can't vote yet, baby!")
  case age < 30:
    fmt.Println("You're still figuring yourself out.")
  case age > 30:
    fmt.Println("You think you've got some stuff down pat.")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("The Weekend is here!")
  default:
    fmt.Println("Time to make that moolah.")
  }

  var a [5]int
  fmt.Println("empty: ", a)
  a[2] = 42
  fmt.Println("something in it: ", a)
  fmt.Println(len(a))


  array := [5]float32{1.1, 2.2, 3.3, 4.4, 5.5}

  fmt.Println("float array: ", array)
}
