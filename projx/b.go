package main

import (
	"fmt"
	"github.com/Rickyc14/proju" /* pkg\windows_amd64\github.com\Rickyc14\proju.a */
)

func main () {
  x := []float64{2.45, 1.67, 3.59, 4.78, 2.98}
  fmt.Println(x)
  fmt.Printf("%.3f \n", proju.Average_these(x))

	fmt.Println("Input integer to get factorial!")
	value := proju.Get_int()
	fmt.Println(proju.Factorial(value))
	fmt.Println("Input x^y: ")
	a := proju.Get_float()
	b := proju.Get_int()
	fmt.Println("Here it is: ", proju.Power(a, b))
}
