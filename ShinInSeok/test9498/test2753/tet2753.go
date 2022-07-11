package test2753

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2753() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	strNum1 := strings.TrimSpace(str)

	year, _ := strconv.Atoi(strNum1)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
