package main
//resolve undefined package "time"
import (
  "fmt"
  "os"
  "time"
)

func main() {
  s, sep := "", ""
  start := time.Nanoseconds()
  for _,args := range os.Args[1:] {
    s += sep + args
    sep = " "
  }
  end := time.Nanoseconds()
  fmt.Println(end-start)
}
