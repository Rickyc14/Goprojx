package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("db.txt")
  if err != nil {
    fmt.Println("An error has occurred: ", err)
    return
  }
  defer file.Close()

  // get the file size
  stat, err := file.Stat()
  if err != nil {
    fmt.Println("An error has occurred: ", err)
    return
  }
  // read the file
  bs := make([]byte, stat.Size())
  _, err = file.Read(bs)
  if err != nil {
    fmt.Println("An error has occurred: ", err)
    return
  }

  str := string(bs)
  fmt.Println(str)
}

/*
A defer statement defers (delays, puts off) the execution of a function until
the surrounding function returns. The deferred call's arguments
 are evaluated immediately, but the function call is not executed
  until the surrounding function returns.
*/
