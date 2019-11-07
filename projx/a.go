package main

import (
	"fmt"
	"github.com/Rickyc14/proju" /* pkg\windows_amd64\github.com\Rickyc14\proju.a */
)

type Grades struct {
	red, mat, lin, hum, nat float64
}

func main() {
	var enem_grades Grades
	values := make([]float64, 5)

	for i := 0; i <= 4; i++ {
		values[i] = proju.Get_float()
	}
	fmt.Println(values)
	fmt.Println(proju.Average_these(values))

	x := proju.Get_float()
	y := proju.Get_int()

	enem_grades.red = x
	enem_grades.mat = x
	enem_grades.lin = x
	enem_grades.hum = x
	enem_grades.nat = x

	fmt.Printf("Hello, new gopher! Here: %.2f | %d. \n", x, y)
	fmt.Println(enem_grades)

}
/*
functions need to start with an upper case letter to be exportable

package main == tells the go tool to produce an executable command
instead of a package object that would be imported by other code.

C:\Users\riccl\go\pkg\windows_amd64\github.com\Rickyc14\proju.a
Above is the location of the 'package object' after we run 'go install'
proju.a is now available to other code by importing from "github.com/Rickyc14/proju"

go install == compile and install packages and dependencies
go build == compile packages and dependencies (use './name_of_exec' to run executable)
go run example.go == compile and run Go program
*/
