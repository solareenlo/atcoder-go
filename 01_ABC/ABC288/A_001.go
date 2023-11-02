package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for n > 0 {
		n--
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
