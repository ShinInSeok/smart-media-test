package test2525

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2525() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	blank := strings.Fields(str)

	numTrim := strings.TrimSpace(blank[0])
	numTrim2 := strings.TrimSpace(blank[1])
	numTrim3 := strings.TrimSpace(str2)

	num, _ := strconv.Atoi(numTrim)
	num2, _ := strconv.Atoi(numTrim2)
	num3, _ := strconv.Atoi(numTrim3)

	input := ((num * 60) + num2) + num3
	hour := input / 60
	minute := input % 60
	if hour/24 >= 1 {
		hour := hour % 24
		fmt.Println(hour, minute)
	} else {
		fmt.Println(hour, minute)
	}

}
