package proju

import (
	"fmt"
	"strconv"
)

func Get_float() float64{
	for {
		var float_string string
		fmt.Print("Enter float: ")
		fmt.Scanln(&float_string)
		float_number, error := strconv.ParseFloat(float_string, 64)
		if error == nil {
			return float_number
		}
		fmt.Println("Please enter a valid float.")
	}
}

func Get_int() int{
	for {
		var int_string string
		fmt.Print("Enter integer: ")
		fmt.Scanln(&int_string)
		int_number, error := strconv.Atoi(int_string)
		if error == nil {
			return int_number
		}
		fmt.Println("Please enter a valid integer.")
	}
}

func Get_floatx(str1 string, str2 string, str3 string) float64{
	for {
		var float_string string
		fmt.Print(str1, str2, str3)
		fmt.Scanln(&float_string)
		float_number, error := strconv.ParseFloat(float_string, 64)
		if error == nil {
			return float_number
		}
		fmt.Println("Invalid input.")
	}
}

func Average_these(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs { /* _ is the index which isn't used; x is the value; xs is the slice/array */
		total += x
	}
	return total/float64(len(xs))
}

func Factorial(x int) int {
  if x == 1 {
    return 1
  }
  return x*Factorial(x-1)
}

func Power(x float64, n int) float64 {
  if n == 0 {
    return 1
  }
  return x*Power(x, n-1)
}
