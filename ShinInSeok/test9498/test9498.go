package test9498

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test9498() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	strNum1 := strings.TrimSpace(str)

	num1, _ := strconv.Atoi(strNum1)

	if 90 <= num1 && num1 <= 100 {
		fmt.Println("A")
	}
	if 80 <= num1 && num1 <= 89 {
		fmt.Println("B")
	}
	if 70 <= num1 && num1 <= 79 {
		fmt.Println("C")
	}
	if 60 <= num1 && num1 <= 69 {
		fmt.Println("D")
	}
	if 0 <= num1 && num1 <= 59 {
		fmt.Println("F")
	}
}
