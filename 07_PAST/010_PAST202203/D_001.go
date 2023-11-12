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

	var m [15]int

	var t, n int
	fmt.Fscan(in, &t, &n)
	for t > 0 {
		t--
		for i := 0; i < n; i++ {
			var a int
			fmt.Fscan(in, &a)
			m[i] = max(m[i], a)
		}
		c := 0
		for i := 0; i < n; i++ {
			c += m[i]
		}
		fmt.Fprintln(out, c)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
