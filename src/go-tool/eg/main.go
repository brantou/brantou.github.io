package main


import "time"

func ExtendWith50000ns(t time.Time) time.Time {
  return t.Add(time.Duration(50000))
}
