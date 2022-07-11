package test10869

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test10869() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	split := strings.Fields(str)

	first := split[0]
	second := split[1]

	num1, _ := strconv.Atoi(first)
	num2, _ := strconv.Atoi(second)

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
}
