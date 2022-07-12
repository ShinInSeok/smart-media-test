package test2742

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2742() {
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		fmt.Fprintln(writer, T-i)
	}
}
