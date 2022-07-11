package test2588

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2588() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	strNum1 := strings.TrimSpace(str)
	strNum2 := strings.TrimSpace(str2)

	num1, _ := strconv.Atoi(strNum1)
	num2, _ := strconv.Atoi(strNum2)

	result3 := num1 * (num2 % 10)
	result4 := num1 * ((num2 - (num2 % 10)) % 100)
	result5 := (num1 * (num2 / 100)) * 100
	result6 := num1 * num2

	fmt.Println(result3)
	fmt.Println(result4 / 10)
	fmt.Println(result5 / 100)
	fmt.Println(result6)
}
