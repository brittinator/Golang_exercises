package main

import (
  "fmt"
  "strconv"
)

// no return for this fxn
func sums(nums...int) {
  // fmt.Print(nums, "=")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

// return type is string
func toString(nums...int) string {
  //needs explicit return
  string := ""

  // for i := 0; i < len(nums); i++ {
  //   // Itoa = converts int to string
  //   // Atoi = converts string to int
  //   n := strconv.Itoa(nums[i])
  //   string += n
  //   if i != len(nums) - 1 {
  //     string += "+"
  //   } else {
  //     string += "="
  //   }
  // }
  // another way to iterate through whole array
  for i, num := range nums {
    n := strconv.Itoa(num)
    string += n
    if i != len(nums) - 1 {
      string += "+"
    } else {
      string += "="
    }
  }
  //range iterate over an array, slice, string, map, or channel.
  return string
}

func main() {
  numbers := []int{10,20,5,15}
  addy := toString(numbers...)
  fmt.Print(addy)
  sums(numbers...)

  fmt.Println("Tada!")

}
