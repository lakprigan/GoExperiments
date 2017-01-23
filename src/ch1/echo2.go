package main

import (
  "fmt"
  "os"
)

func main() {
  for index, value := range os.Args[0:] {
    fmt.Print(index)
    fmt.Print(":")
    fmt.Println(value)
  }
}
