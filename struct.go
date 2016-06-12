/*
Go’s structs are typed collections of fields.
They’re useful for grouping data together to form records
*/

package main

import "fmt"


type wizard struct{
  name string
  house string
  good bool
}

func main(){
  hp := wizard{
    name: "Harry Potter",
    house: "Gryffindor",
    good: true,
  }

  tr := wizard {
    name: "Tom Riddle",
    house: "Slytherin",
    good: true,
}
  fmt.Println(hp, tr)

  fmt.Println("House ", hp.house)

  // to reassign things, need to get the pointer for the object first
  trPointer := &tr //& prefix gets you the pointer to the struct
  trPointer.name = "Voldemort"
  trPointer.good = false

  fmt.Println(tr)

}
