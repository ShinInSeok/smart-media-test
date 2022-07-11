package test14681

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test14681() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	numTrim := strings.TrimSpace(str)
	numTrim2 := strings.TrimSpace(str2)

	num, _ := strconv.Atoi(numTrim)
	num2, _ := strconv.Atoi(numTrim2)

	if num > 0 && num2 > 0 {
		fmt.Println(1)
	} else if num < 0 && num2 > 0 {
		fmt.Println(2)
	} else if num < 0 && num2 < 0 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
