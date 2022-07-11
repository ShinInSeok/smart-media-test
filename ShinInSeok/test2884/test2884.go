package test2884

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2884() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	split := strings.Fields(str)

	first := split[0]
	second := split[1]

	H, _ := strconv.Atoi(first)
	M, _ := strconv.Atoi(second)

	if M < 45 {
		H--
		M = 60 - (45 - M)
		if H < 0 {
			H = 23
		}
		fmt.Println(H, M)
	} else {
		fmt.Println(H, (M - 45))
	}
}
