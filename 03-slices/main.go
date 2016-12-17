package main

import "fmt"

func main() {
  // direct creation
  var ints []int = []int{1,2,3,4,5}
  fmt.Printf("len=%d cap=%d\n", len(ints), cap(ints), ints)
  // indirect creation
  var bytes []byte = make([]byte, 4, 4)
  bytes[0] = 0x21
  bytes[1] = 0x23
  fmt.Printf("len=%d cap=%d %#v\n", len(bytes), cap(bytes), bytes)
  // access to slice elements
  fmt.Printf("%#v\n", bytes[2:]) // from 2 index to end
  fmt.Printf("%#v\n", bytes[:2]) // until 2 index
  // growing slices
  ints = append(ints, 12)
  fmt.Printf("len=%d cap=%d %#v\n", len(ints), cap(ints), ints)
  // concatenating slices
  bytes = append(bytes, []byte{0xf5, 0xf1, 0xaa}...)
  fmt.Printf("len=%d cap=%d %#v\n", len(bytes), cap(bytes), bytes)
}
