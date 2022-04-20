package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	c := make([]int, n+1)
	mi := [2]int{1 << 60, 1 << 60}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
		mi[i%2] = min(mi[i%2], c[i])
	}

	var q int
	fmt.Fscan(in, &q)
	ans := 0
	ac := [2]int{}
	for i := 0; i < q; i++ {
		var t, a int
		fmt.Fscan(in, &t, &a)
		if t == 1 {
			var b int
			fmt.Fscan(in, &b)
			if c[a]-ac[a%2] >= b {
				ans += b
				c[a] -= b
				mi[a%2] = min(mi[a%2], c[a]-ac[a%2])
			}
		} else if t == 2 {
			if mi[1] >= a {
				ans += a * ((n + 1) / 2)
				ac[1] += a
				mi[1] -= a
			}
		} else if t == 3 {
			if mi[0] >= a && mi[1] >= a {
				ans += a * n
				ac[0] += a
				ac[1] += a
				mi[0] -= a
				mi[1] -= a
			}
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
