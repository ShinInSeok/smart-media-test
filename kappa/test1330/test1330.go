package test1330

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test1330() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	blank := strings.Fields(str)

	num1, _ := strconv.Atoi(blank[0])
	num2, _ := strconv.Atoi(blank[1])

	if num1 > num2 {
		fmt.Println(">")
	} else if num1 < num2 {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}
