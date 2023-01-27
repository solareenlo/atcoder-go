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

	var n, x, y, z int
	fmt.Fscan(in, &n, &x, &y, &z)

	const N = 1001
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	c := make([]int, N)
	for j := 1; j <= x; j++ {
		d, e := -1, 0
		for i := 1; i <= n; i++ {
			if a[i] > d && c[i] == 0 {
				e = i
				d = a[i]
			}
		}
		c[e] = 1
	}

	for j := 1; j <= y; j++ {
		d, e := -1, 0
		for i := 1; i <= n; i++ {
			if b[i] > d && c[i] == 0 {
				e = i
				d = b[i]
			}
		}
		c[e] = 1
	}

	for j := 1; j <= z; j++ {
		d, e := -1, 0
		for i := 1; i <= n; i++ {
			if a[i]+b[i] > d && c[i] == 0 {
				e = i
				d = a[i] + b[i]
			}
		}
		c[e] = 1
	}

	for i := 1; i <= n; i++ {
		if c[i] == 1 {
			fmt.Fprintln(out, i)
		}
	}
}
