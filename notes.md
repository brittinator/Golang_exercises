# Golang notes


Go is a compiled, statically typed language, that is relatively quick to compile to machine code. It also has garbage collection.


generic package is main

```
package main

import "fmt"

func main() {
  fmt.Print("Hello there")
}
```

Basic Types:
```
bool
string
int
byte
rune
float32 float 64
```

Constants are defined like `const NumOfPlanets = 8`

## For
For loop is similar to javascript or java, with init, condition, and post statement
init and post statements are optional

```
for i := 0; i < 10; i++ {
  do something
}
```

for is also Go's 'while'

```
treats = 5
for treats > 0 {
  fmt.Println("YUM")
  treats -= treats
}
```
an infinite loop would look like `for {}`

## If/Else and Switch

```
n := 8
if n > 6 {
  so something
} else if n < 6 {
  do something else
} else {
  something else here
}
```

```
package main

import "fmt"

func main() {
  fmt.Print("Hello there")
}
```

## pointers


* Go has pointers. A pointer (like *this) holds the memory address of a variable.
* (*this) is a pointer to a this value
* to generate a pointer, use &

```
name := "Britt"
p := &name  <-- p is a pointer to the name value

passing pointers
func PointersOhMy(*p1 *p2) { do something}
```


## structs

Go is not really an OO language, so if you want to pass around
something with fields you use `struct`, which is a collection of fields

* access fields trhough struct pointer

```
type Costume struct {
  hat bool
  length int
  color string
  makeup string
}

func main() {
  c := Costume{false, 100, "red", "eye shadow, lipstick, cover up"}
  pointer := &c
  pointer.hat // false
  pointer.color // "red"
  anotherCostume := {length: 500}
  // everything else zero value,
  // hat = false, color = "", makeup = ""
}
```

## arrays

* array has fixed size
* but a slice is dynamically sized view of the elements in an array
* use slices to deal with arrays
* a slice is an abstraction built on top of Go's array type
blog post about slices: https://blog.golang.org/go-slices-usage-and-internals
