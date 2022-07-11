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

	blank := strings.Fields(str)

	numTrim := strings.TrimSpace(blank[0])
	numTrim2 := strings.TrimSpace(blank[1])

	num, _ := strconv.Atoi(numTrim)
	num2, _ := strconv.Atoi(numTrim2)

	if num == 0 && num2 < 45 {
		num := 24
		time := (((num * 60) + num2) - 45)
		hour := time / 60
		minute := time % 60
		fmt.Println(hour, minute)

	} else {
		time := (((num * 60) + num2) - 45)
		hour := time / 60
		minute := time % 60
		fmt.Println(hour, minute)
	}
}
