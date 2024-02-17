package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	x := make([]int, m)
	y := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	sort.Ints(x)
	sort.Ints(y)
	l, r := 0, int(1e9)
	ans := 0
	for l <= r {
		md := (l + r) / 2
		c := 2
		for i, j := 0, 0; i < m; i++ {
			if x[i]-x[j] > md {
				c++
				j = i
			}
		}
		for i, j := 0, 0; i < m; i++ {
			if y[i]-y[j] > md {
				c++
				j = i
			}
		}
		if c <= n {
			ans = md
			r = md - 1
		} else {
			l = md + 1
		}
	}
	fmt.Println(ans)
}
