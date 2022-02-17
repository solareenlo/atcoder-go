package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 114514

var (
	c   string
	e   = make([][]int, N)
	s   int
	f   = make([]int, N)
	mid int
	u   int
)

func dfs(x, fa int) int {
	rs := 0
	ty := 0
	if fa > 0 {
		ty = 1
	}
	if c[x] == 'W' {
		rs = 1
	}
	for _, y := range e[x] {
		if y != fa && dfs(y, x) != 0 {
			rs = 1
			s += 2
			ty++
		}
	}
	if rs != 0 {
		if (ty&1 != 0) != (c[x] == 'W') {
			f[x] = 2
			s++
		}
	} else {
		f[x] = -1
	}
	return rs
}

func sol(x, fa, d int) {
	d += f[x]
	if d >= mid {
		mid = d
		u = x
	}
	for _, y := range e[x] {
		if y != fa && f[y] != -1 {
			sol(y, x, d)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}

	fmt.Fscan(in, &c)
	c = " " + c

	var i int
	for i = 1; i <= n; i++ {
		if c[i] == 'W' {
			break
		}
	}
	if i > n {
		fmt.Println(0)
		return
	}
	dfs(i, 0)
	mid = -1 << 60
	sol(i, 0, 0)
	mid = -1 << 60
	sol(u, 0, 0)
	if s > 1 {
		fmt.Println(s - mid)
	} else {
		fmt.Println(1)
	}
}
