package main


import (
  "fmt"
  "time"
)

func main() {
  test_len := 10000000
  start := time.Now()
  s := make([]int, test_len, test_len)
  for i := 0; i < test_len; i++ {
    s = append(s, i)
  }
  fmt.Println(time.Now().Sub(start).String())

  start = time.Now()
  s1 := make([]int, test_len)
  for i := 0; i < test_len; i++ {
    s1[i] = i
  }
  fmt.Println(time.Now().Sub(start).String())
}
