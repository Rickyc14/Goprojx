package main

import (
  "fmt"
  "github.com/Rickyc14/proju"
)

func main() {
  grades := make(map[string]float64)
  weights := make(map[string]float64)
  topics := [5]string{"red", "mat", "lin", "hum", "nat"}

  for i := 0; i < 5; i++ {
    for {
      grades[topics[i]] = proju.Get_floatx("Please enter ", topics[i], " grade: ")
      if grades[topics[i]] >= 0 && grades[topics[i]] <= 1000 {
        break
      }
      fmt.Println("Invalid grade.")
    }
  }

  fmt.Print("\n")
  e := 0
  var total_weight float64

  for {
    if e == 5 && total_weight != 10 {
      fmt.Println("Invalid sum of weights. Sum must be 10.")
      e = 0
      total_weight = 0
    } else if e == 5 && total_weight == 10 {
      break
    }
    for {
      weights[topics[e]] = proju.Get_floatx("Please enter ", topics[e], " weight: ")
      if weights[topics[e]] >= 0 && weights[topics[e]] <= 10 {
        total_weight += weights[topics[e]]
        break
      }
      fmt.Println("Invalid weight.")
    }
    e += 1
  }

  fmt.Print("\n")
  var total float64
  total = grades["red"] * weights["red"]
  total += grades["mat"] * weights["mat"]
  total += grades["lin"] * weights["lin"]
  total += grades["hum"] * weights["hum"]
  total += grades["nat"] * weights["nat"]

  fmt.Printf("\nYour final grade is: %.3f", total/float64(10))
}
