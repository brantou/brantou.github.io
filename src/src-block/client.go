package main


import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
  "strconv"
  "time"
)

func readData(bfRd io.Reader, dataChan chan []byte, endChan chan error) {
  for {
    buf := make([]byte, 100)
    n, err := bfRd.Read(buf)
    if err != nil {
      endChan <- err
      return
    }
    fmt.Printf("read data %s size \n", strconv.Itoa(n))
    dataChan <- buf
  }
}

func main() {
  file, err := os.Open("client.go")
  if err != nil {
    log.Fatalln(err)
  }

  var data []byte
  bfRd := bufio.NewReader(file)
  var dataChan = make(chan []byte)
  var endChan = make(chan error)
  go readData(bfRd, dataChan, endChan)

  for {
    var err error
    var sub_data []byte
    select {
    case err = <-endChan:
      fmt.Println(err)
      break
    case sub_data = <-dataChan:
      data = append(data, sub_data...)
    case <-time.After(2 * time.Second):
      fmt.Println("timeout")
      break
    }
  }

  fmt.Println(string(data))
}
