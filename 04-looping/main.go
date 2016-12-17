package main

import "fmt"

func main() {
  // simple for loop
  for i := 1; i <= 3; i++ {
    fmt.Printf("index=%v\n", i)
  }
  // the loop can be done with any iterable operation
  letters := []rune{'a', 'b', 'c'}
  for ;len(letters) > 0; {
    fmt.Printf("%#v\n", letters)
    letters = letters[:len(letters)-1]
  }
  // range loop
  bytes := []byte{0xf1, 0xf2, 0xf3}
  for i := range bytes {
    fmt.Printf("index=%d, %#v, %#v\n", i, bytes[i], bytes)
  }
  // range loop, using value
  for i, value := range bytes {
    fmt.Printf("index=%d, %#v, %#v\n", i, value, bytes)
  }
  // range loop, using value only
  for _, value := range bytes {
    fmt.Printf("%#v, %#v\n", value, bytes)
  }
  // forever loop
  times := 0
  for {
    fmt.Println("I will never stop")
    if times == 3 {
      break
    }
    times = times + 1
  }
  fmt.Println("Unless you break it...")
}
