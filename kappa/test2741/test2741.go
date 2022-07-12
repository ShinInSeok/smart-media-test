package test2741

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2741() {
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)
	for i := 1; i <= T; i++ {
		fmt.Fprintln(writer, i)
	}
}
