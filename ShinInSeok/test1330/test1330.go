package test1330

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test1330() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	split := strings.Fields(str)

	first := split[0]
	second := split[1]

	num1, _ := strconv.Atoi(first)
	num2, _ := strconv.Atoi(second)

	if num1 > num2 {
		fmt.Println(">")
	}
	if num1 < num2 {
		fmt.Println("<")
	}
	if num1 == num2 {
		fmt.Println("==")
	}
}
