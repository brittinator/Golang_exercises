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

## testing

```
import "testing"

func TestIsEqual(t *testing.T) {

}
```

## interface

* defined set of methods. A value of interface type can hold any value that implements those methods.

example of the namer interface, which is implemented by `Customer` and `User`:

    ```
    package main

    import (
      "fmt"
    )

    type User struct {
      FirstName, LastName string
    }

    func (u *User) Name() string {
      return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
    }

    type Customer struct {
      Id       int
      FullName string
    }

    func (c *Customer) Name() string {
      return c.FullName
    }

    type Namer interface {
      Name() string
    }

    func Greet(n Namer) string {
      return fmt.Sprintf("Dear %s", n.Name())
    }

    func main() {
      u := &User{"Matt", "Aimonetti"}
      fmt.Println(Greet(u))
      c := &Customer{42, "Francesc"}
      fmt.Println(Greet(c))
    }
    ```

example of an animal interface; in this case an animal is anything that can speak.

  ```
  type Animal interface {
    Speak() string
  }
  ```

Implementations of animal type must have the speak function.
  ```
  type Cat struct {}

  func Speak(c Cat) string {
    return "Meow!"
  }

  type Elephant struct{}

  func Speak(e Elephant) string {
    return "Trumpet Sound"
  }
  ... etc
  func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
  ```

? Why are the above method signatures `func Speak(c Cat) {}` and not `func Speak(c *Cat){}`?
Because down below, you are passing Cat{} into the animal slice, but if you pass `&Cat{}` or `new(Cat)` it will work then

> prog.go:40: cannot use Cat literal (type Cat) as type Animal in array element:
>    Cat does not implement Animal (Speak method requires pointer receiver)
> This error message is a bit confusing at first, to be honest. What it’s saying is not that the interface Animal demands that you define your method as a pointer receiver, but that you have tried to convert a Cat struct into an Animal interface value, but only *Cat satisfies that interface. You can fix this bug by passing in a *Cat pointer to the Animal slice instead of a Cat value, by using new(Cat) instead of Cat{} (you could also say &Cat{}, I simply prefer the look of new(Cat)):

> Let’s go in the opposite direction: let’s pass in a *Dog pointer instead of a Dog value, but this time we won’t change the definition of the Dog type’s Speak method:

`animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}`
> This also works (http://play.golang.org/p/UZ618qbPkj), but recognize a subtle difference:  we didn’t need to change the type of the receiver of the Speak method. This works because a pointer type can access the methods of its associated value type, but not vice versa. That is, a *Dog value can utilize the Speak method defined on Dog, but as we saw earlier, a Cat value cannot access the Speak method defined on *Cat.





Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface.
That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value.
An interface value is constructed of two words of data; one word is used to point to a method table for the value’s underlying type,
and the other word is used to point to the actual data being held by that value.