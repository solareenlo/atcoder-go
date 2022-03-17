package main

import (
	"bufio"
	"fmt"
	"os"
)

var e = make([][]int, 2020)

func t(x, p, d, m int) int {
	if d == m {
		return 1
	}
	s := 1
	for _, i := range e[x] {
		if i != p {
			s += t(i, x, d+1, m)
		}
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u-1] = append(e[u-1], v-1)
		e[v-1] = append(e[v-1], u-1)
	}

	a := 0
	if k&1 != 0 {
		for i := 0; i < n; i++ {
			for _, j := range e[i] {
				a = max(a, t(i, j, 0, k/2)+t(j, i, 0, k/2))
			}
		}
	} else {
		for i := 0; i < n; i++ {
			a = max(a, t(i, -1, 0, k/2))
		}
	}

	fmt.Println(n - a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
