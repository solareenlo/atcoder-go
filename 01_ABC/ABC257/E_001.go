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

	c := make([]int, 10)
	m := 1_000_000_000
	for i := 1; i <= 9; i++ {
		fmt.Fscan(in, &c[i])
		m = min(m, c[i])
	}

	l := n / m
	for i := l; i >= 1; i-- {
		for j := 9; j >= 1; j-- {
			if n-c[j] >= (i-1)*m {
				n -= c[j]
				fmt.Fprint(out, j)
				break
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
