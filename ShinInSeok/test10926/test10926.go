package test10926

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test10926() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	split := strings.Fields(str)

	first := split[0]

	fmt.Print(first + "??!")
}
