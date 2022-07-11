package test1000to1008

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test1008() {

	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	split := strings.Fields(str)

	first := split[0]
	second := split[1]

	num1, _ := strconv.Atoi(first)
	num2, _ := strconv.Atoi(second)

	no1 := float64(num1)
	no2 := float64(num2)

	fmt.Println(no1 / no2)
}
