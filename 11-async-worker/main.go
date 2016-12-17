package main

import (
  "fmt"
  "time"
  "math/rand"
)

var r = rand.New(rand.NewSource(99))

type Message struct {
  command string
  aggregate int
}

type Worker struct {
  work func (Message, *Worker)
}

type Supervisor struct {
  workers []*Worker
}
type ISupervisor interface {
  CreateWorker(func (Message, *Worker))
  Handle(Message)
}

func NewSupervisor() Supervisor {
  s := Supervisor{
    workers: []*Worker{},
  }
  s.CreateWorker(func (msg Message, worker *Worker) {
     fmt.Printf(
       "%#v worker, received message %#v\n",
       (*worker), msg,
    )
    fmt.Printf("Starting heavy work for message %#v.\n", msg)
    time.Sleep(time.Duration(r.Intn(20)) * time.Second)
    fmt.Printf("Heavy work for message %#v ended.\n", msg)
    ch <- 1
  })
  return s
}

func (s *Supervisor) CreateWorker(work func (Message, *Worker)) {
  s.workers = append(s.workers, &Worker{
    work: work,
  })
}

func (s *Supervisor) Handle(msg Message) {
  for _, worker := range s.workers {
    go worker.work(msg, worker)
  }
}

var msgs = []Message{
  Message {
    command: "RegisterCustomer",
    aggregate: 86346,
  },
  Message {
    command: "AddAddressForCustomer",
    aggregate: 65236,
  },
  Message {
    command: "SendWelcomeEmail",
    aggregate: 61531,
  },
  Message {
    command: "UpdateCustomerProfile",
    aggregate: 1876122,
  },
}

var pending = 0
var ch = make(chan int)
var quit = make(chan bool)

func main() {
  sup := NewSupervisor()
  for _, msg := range msgs {
    sup.Handle(msg)
    time.Sleep(1 * time.Second)
    pending = pending + 1
  }
  go func () {
      for {
        select {
        case  _ = <- ch:
            pending = pending - 1
            if pending == 0 {
              quit <- true
            }
        }
      }
  } ()
  for {
    if <-quit {
      fmt.Println("All done, quitting.")
      break
    }
  }
}
