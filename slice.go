package main

import "fmt"

func main()  {
  mySlice := make([]string, 3)
  fmt.Println(mySlice)

  mySlice[0] = "raspberry"
  mySlice[1] = "strawberry"
  mySlice[2] = "marionberry"

  mySlice = append(mySlice, "blackberry", "cranberry")
  fmt.Println(mySlice)

  sliver := mySlice[2:4] //up to but not including 4th index
  fmt.Println(sliver)
}
