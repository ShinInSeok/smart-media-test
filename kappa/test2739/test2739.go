package test2739

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2739() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	gugudan := strings.TrimSpace(str)

	num, _ := strconv.Atoi(gugudan)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}
