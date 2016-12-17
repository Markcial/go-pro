package main

import "fmt"

type Message struct {
  command string
  aggregate int
}

func Handle(channel chan Message, close chan bool) {
  for {
    select {
    case msg := <- channel:
        fmt.Printf(
          "%s for order id %d received.\n",
          msg.command,
          msg.aggregate,
        )
      case <-close:
        fmt.Println("quit")
			  return
    }
  }
}

func main() {
  // basic usage
  var channel = make(chan string)
  go func() {
    fmt.Printf("%s\n", <-channel)
    fmt.Printf("%s\n", <-channel)
  }()
  channel <- "hola"
  channel <- "mundo"

  // message example
  var msg = Message{
    command: "DeliveryStarted",
    aggregate: 9281,
  }
  ch := make(chan Message)
  cl := make(chan bool)
  go Handle(ch, cl)
  ch <- msg
  ch <- Message{
    command: "DeliveryFinished",
    aggregate: 72826,
  }
  cl <- true
}
