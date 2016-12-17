package main

import (
  "fmt"
  "time"
  "sync"
  "math/rand"
  "./supervisor"
  "./message"
)

var r = rand.New(rand.NewSource(99))
var wg sync.WaitGroup
var msgs = message.NewCollection(map[string]int{
  "RegisterCustomer": 86346,
  "AddAddressForCustomer": 65236,
  "SendWelcomeEmail": 61531,
  "UpdateCustomerProfile": 1876122,
})

func main() {
  sup := supervisor.New()
  sup.CreateWorker(func (msg message.Message, worker *supervisor.Worker) {
    fmt.Printf(
       "%#v worker, received message %#v\n",
       (*worker), msg,
    )
    fmt.Printf("Starting heavy work for message %#v.\n", msg)
    time.Sleep(time.Duration(r.Intn(20)) * time.Second)
    fmt.Printf("Heavy work for message %#v ended.\n", msg)
    wg.Done()
  })
  for _, msg := range msgs {
    wg.Add(1)
    sup.Handle(msg)
    time.Sleep(1 * time.Second)
  }

  wg.Wait()
  fmt.Println("All done, quitting.")
}
