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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var l, a int
		fmt.Fscan(in, &l, &a)
		fmt.Fprintln(out, max(0, l-calc(a)+1))
	}
}

func calc(a int) int {
	if a == 0 {
		return 1
	}
	return int(2.5*float64(a) - 0.5)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
