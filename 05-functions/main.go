package main

import "fmt"
import "errors"

func variadicReception(bytes ...byte) {
  fmt.Printf("%#v\n", bytes)
}

func returnDefinition(a int32, b int32) (int32, error) {
  if a < 0||b < 0  {
    return 0, errors.New("Numbers must be greater than 0")
  }

  return a+b, nil
}

func returnDeclaration(a int32, b int32) (sum int32, err error) {
  if a < 0||b < 0 {
    return 0, errors.New("Numbers must be greater than 0")
  }
  sum = a + b
  return
}

func main() {
  // functions can be assigned as lambdas
  who := "Jane"
  fn := func () {
    // the outer scope is still reachable here
    fmt.Printf("Hi %s\n", who)
  }
  fn()
  who = "Mary"
  fn()
  // called directly on creation
  func (animal string) {
    fmt.Printf("I love %s\n", animal)
  } ("cats")

  // using variadics
  variadicReception(byte(0xaa), byte(0xf1), byte(0x31))
  variadicReception([]byte{0xaa, 0xf1, 0x31}...)

  // calling functions
  sum, err := returnDefinition(3, 6)
  fmt.Printf("%#v\n", sum)
  fmt.Printf("%#v\n", err)
  sum, err = returnDeclaration(-4, 5)
  fmt.Printf("%#v\n", sum)
  fmt.Printf("%#v\n", err)
}
