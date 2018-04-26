package main


import "fmt"

type Outline struct {
  Label string
  Type  string
}

func (ol *Outline) Label() string {
  return ol.Label
}

const (
  name = "brantou"
)

const email = "brantou89@gmail.com"

var (
  cmd  string
  args string
)

func main() {
  fmt.Printf("hello world!")
}
