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

Go has pointers. A pointer (like *this) holds the memory address of a variable.
*this is a pointer to a this value
to generate a pointer, use &