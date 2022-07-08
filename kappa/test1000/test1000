package test1000

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test1000() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	blank := strings.Fields(str)

	numStr1 := blank[0]
	numStr2 := blank[1]

	num1, _ := strconv.Atoi(numStr1)
	num2, _ := strconv.Atoi(numStr2)

	fmt.Println(num1 + num2)
}
