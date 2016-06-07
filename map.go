// Maps are key:value pairs in Go
package main

import "fmt"

func main()  {
  northwestStates := make(map[string]int)
  northwestStates["Washington"] = 24
  northwestStates["Oregon"] = 5
  northwestStates["Kentucky"] = 100

  fmt.Println("the map: ", northwestStates)

  fmt.Println(northwestStates["Oregon"])

  delete(northwestStates, "Kentucky")
  fmt.Println(northwestStates)
  _, kentucky := northwestStates["Kentucky"]
  fmt.Println("Is Kentucky here? ", kentucky)

  // initialize and declare at the same time
  anotherMap := map[string]string{
    "kittens": "adorable",
    "cats": "cute",
    "dogs": "okay",
    "kangaroo": "neat"}

  fmt.Println(anotherMap, "length: ", len(anotherMap))
}
