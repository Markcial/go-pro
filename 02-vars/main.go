package main

import "fmt"

// package scope
// constants
// block declaration
const (
  PI = 3.14
  ENABLED = true
)
// single declaration
const TYPE = "type"

// var declaration
var (
  integer int32 = 12
  letter []rune = []rune{'ñ'}
  bytes []byte = []byte{'ß'}
)

func main() {
  // function scope
  // explicit
  var hello string = "world"
  // implicit
  whoami := []interface{}{12, string("what"), 5.4444}

  fmt.Printf("%#v\n", PI)
  fmt.Printf("%#v\n", ENABLED)
  fmt.Printf("%#v\n", TYPE)

  fmt.Printf("%#v\n", integer)
  fmt.Printf("%#v\n", letter)
  fmt.Printf("%#v\n", bytes)

  fmt.Printf("%#v\n", hello)
  fmt.Printf("%#v\n", whoami)
}
