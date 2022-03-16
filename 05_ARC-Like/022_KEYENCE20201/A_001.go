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
	c := make([]int, n)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	d := make([]int, n)
	for i := range d {
		fmt.Fscan(in, &d[i])
	}

	a := 0
	ans := 0
	for i := 0; i < n; i++ {
		a = max(a, c[i])
		ans = max(ans, a*d[i])
		fmt.Fprintln(out, ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
