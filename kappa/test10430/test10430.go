package test10430

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test10430() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	blank := strings.Fields(str)

	numStr1 := blank[0]
	numStr2 := blank[1]
	numStr3 := blank[2]

	num1, _ := strconv.Atoi(numStr1)
	num2, _ := strconv.Atoi(numStr2)
	num3, _ := strconv.Atoi(numStr3)

	fmt.Println((num1 + num2) % num3)
	fmt.Println(((num1 % num3) + (num2 % num3)) % num3)
	fmt.Println((num1 * num2) % num3)
	fmt.Println(((num1 % num3) * (num2 % num3)) % num3)
}
