package main

import (
  "fmt"
  "net/http"
  "html/template"
)

type WebPage struct {
  Title string
  Header string
  Table [][]string
}

func handler(w http.ResponseWriter, r *http.Request) {
  p := &WebPage{
    Title: "My Webpage",
    Header: "Webpage with table data",
    Table: [][]string{
      []string{"a", "row", "of", "data"},
      []string{"another", "row", "of", "data"},
    },
  }
  t, _ := template.ParseFiles("./10-webserver/index.html")
  t.Execute(w, p)
}

func main() {
  fmt.Println("Starting Webserver on http://localhost:8080/")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
