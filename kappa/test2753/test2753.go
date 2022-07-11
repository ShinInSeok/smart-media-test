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

	numTrim := strings.TrimSpace(str)

	num, _ := strconv.Atoi(numTrim)

	if (((num % 4) == 0) && ((num % 100) != 0)) || ((num % 400) == 0) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
