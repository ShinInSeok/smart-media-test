package test18108

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test18108() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	strNum := strings.TrimSpace(str)

	num1, _ := strconv.Atoi(strNum)

	fmt.Println(num1 - 543)
}
