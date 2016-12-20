package main

import (
  "fmt"
  "time"
  "math/rand"
  "./supervisor"
)

var r = rand.New(rand.NewSource(99))
var msgs = supervisor.NewMessageCollection(map[string]int{
  "RegisterCustomer": 86346,
  "AddAddressForCustomer": 65236,
  "SendWelcomeEmail": 61531,
  "UpdateCustomerProfile": 1876122,
})

func main() {
  sup := supervisor.New()
  sup.CreateWorker(func (msg supervisor.Message) {
    fmt.Printf("Received message %#v\n", msg)
    fmt.Printf("Starting heavy work for message %#v.\n", msg)
    time.Sleep(time.Duration(r.Intn(20)) * time.Second)
    fmt.Printf("Heavy work for message %#v ended.\n", msg)
  })
  for _, msg := range msgs {
    sup.Handle(msg)
  }
  sup.Wait()
  fmt.Println("All done, quitting.")
}
