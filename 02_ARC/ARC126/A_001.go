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

	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		s := min(b/2, c+a/2)
		c -= s
		if c < 0 {
			a += c * 2
			c = 0
		}
		z := min(a, (c*2+a)/5)
		fmt.Fprintln(out, z+s)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
