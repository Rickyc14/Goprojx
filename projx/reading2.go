package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  bs, err := ioutil.ReadFile("db.txt")
  if err != nil {
    fmt.Println("An error has occurred: ", err)
    return
  }
  str := string(bs)
  fmt.Println(str)
}
