package test9498

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test9498() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	numTrim := strings.TrimSpace(str)

	num, _ := strconv.Atoi(numTrim)

	if num >= 90 && num <= 100 {
		fmt.Println("A")
	} else if num >= 80 && num <= 89 {
		fmt.Println("B")
	} else if num >= 70 && num <= 79 {
		fmt.Println("C")
	} else if num >= 60 && num <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
