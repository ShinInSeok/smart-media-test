package test2480

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2480() {

	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	//str2, _ := br.ReadString(('\n'))

	split := strings.Fields(str)
	// split2 := strings.Fields(str2)

	first := split[0]
	second := split[1]
	third := split[2]

	a, _ := strconv.Atoi(first)
	b, _ := strconv.Atoi(second)
	c, _ := strconv.Atoi(third)

	if a == b && b == c && a == c {
		fmt.Println(10000 + a*1000)
	}
	if a == b && b != c {
		fmt.Println(1000 + a*100)
	}
	if a == c && a != b {
		fmt.Println(1000 + c*100)
	}
	if b == c && a != c {
		fmt.Println(1000 + b*100)
	}
	if a > b && a > c && c != a && c != b && a != b {
		fmt.Println(a * 100)
	}
	if a < b && b > c && a != c && a != b && c != b {
		fmt.Println(b * 100)
	}
	if a < c && b < c && c != a && c != b && a != b {
		fmt.Println(c * 100)
	}

}
