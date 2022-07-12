package test10950

import (
	"fmt"
)

func Test10950() {
	var t int
	fmt.Scan(&t)

	var num1 int
	var num2 int

	for i := 1; i <= t; i++ {
		fmt.Scan(&num1, &num2)
		fmt.Println(num1 + num2)
	}

}
