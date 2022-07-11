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

	split := strings.Fields(str)
	split2 := strings.Fields(str2)

	first := split[0]
	second := split[1]
	third := split2[0]

	H, _ := strconv.Atoi(first)
	M, _ := strconv.Atoi(second)
	M2, _ := strconv.Atoi(third)

	if M+M2 < 60 {
		fmt.Println(H, M+M2)
	} else {
		hour := (M + M2) / 60
		minute := (M + M2) % 60
		if H+hour < 24 {
			fmt.Println(H+hour, minute)
		} else {
			fmt.Println(H+hour-24, minute)
		}
	}
}
