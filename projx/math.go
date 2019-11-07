package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(sin(math.Pi/2))
}

func sin(x float64) float64 {
  var result float64
  counter := 1
  for i := 1; i <= 13; i += 2 {
    if counter % 2 != 0 {
      result += power(x, i)/float64(factorial(i))
    } else {
      result -= power(x, i)/float64(factorial(i))
    }
    counter += 1
  }
  return result
}

func factorial(x int) int {
  if x == 1 {
    return 1
  }
  return x*factorial(x-1)
}

func power(x float64, n int) float64 {
  if n == 0 {
    return 1
  }
  return x*power(x, n-1)
}
