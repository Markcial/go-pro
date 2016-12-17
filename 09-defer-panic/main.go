package main

import "fmt"

var dataCorrupted = false

func SomeDangerousOperation() {
  defer Rollback()
  fmt.Println("Writting to the database.")
  if dataCorrupted {
    panic("Unable to complete operation")
  }
}

func Rollback() {
  if r := recover(); r != nil {
    fmt.Println("Whoopsie! Rolling back!", r)
  } else {
    fmt.Println("Everything went just fine.")
  }
}

func main() {
  SomeDangerousOperation()
  dataCorrupted = true
  SomeDangerousOperation()
}
