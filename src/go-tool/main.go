package main


import "fmt"

type Outline struct {
  Label string
  Type  string
}

func (ol *Outline) Label() string {
  return ol.Label
}

var name string

var (
  cmd  string
  args string
)

func main() {
  fmt.Printf("hello world!")
}
