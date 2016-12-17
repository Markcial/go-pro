package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  // blocking calls
  for i := 0; i < 3; i++ {
    fmt.Printf("%d times...\n", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("%d done\n", i)
  }
  // async goroutine
  var wg sync.WaitGroup
  for i := 0; i < 3; i++ {
    wg.Add(1)
    go func (i int) {
      fmt.Printf("%d times...\n", i)
      time.Sleep(2 * time.Second)
      fmt.Printf("%d done\n", i)
      wg.Done()
    }(i)
  }
  wg.Wait()
}
