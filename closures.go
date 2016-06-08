package main

import (
  "fmt"
)

func intSeq() func() int {
  i := 0
  // return value of intSeq() method is a closure, an anonymous fxn
  return func() int {
    if i == 0 {
      i += 1
    } else {
      i += i
    }
    return i
  }
}

func main() {
  myNextInt := intSeq()
  fmt.Println(myNextInt())
  fmt.Println(myNextInt())
  fmt.Println(myNextInt())
  fmt.Println(myNextInt())

  initialized := intSeq()
  fmt.Println(initialized())

}
