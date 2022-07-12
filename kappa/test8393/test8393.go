package test8393

import (
	"fmt"
)

func Test8393() {
	var t int
	fmt.Scan(&t)

	sum := 0
	for i := 1; i <= t; i++ {
		sum += i

	}
	fmt.Println(sum)

}
