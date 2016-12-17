package main

import "fmt"

// alias
type Relation map[string]int

type Email string
// struct
type Order struct {
  id int32
  products []string
}
type Customer struct {
  email Email
  id int32
  orders []Order
}
type Client interface {
  AddOrder(orders... Order)
  OrdersCount() int
}

func (c *Customer) AddOrder(orders... Order) {
  c.orders = append(c.orders, orders...)
}

func (c *Customer) OrdersCount() int {
  return len(c.orders)
}

func main() {
  // direct notation
  relation := Relation{
    "eggs":5,
    "spam":8,
  }
  fmt.Printf("%#v\n", relation)
  // indirect notation
  relation = make(Relation, 2)
  relation["foo"] = 1
  relation["bar"] = 2
  fmt.Printf("%#v\n", relation)

  customer := &Customer{}
  fmt.Printf("%#v\n", customer)
  order := Order{
    id: 3214,
    products: []string{"milk", "soy", "eggs"},
  }
  customer.AddOrder(order)
  fmt.Printf("%#v\n", customer)
  fmt.Printf("%d orders\n", customer.OrdersCount())
}
