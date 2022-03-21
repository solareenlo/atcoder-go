package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 233333

var (
	k = make([]int, N)
	e = make([][]int, N)
	f int
)

func dfs(x, fa int) {
	k[x] = k[fa] ^ 2 | 1
	for _, y := range e[x] {
		if k[y] != 0 {
			if k[x] == k[y] {
				f |= 1
			} else {
				f |= 0
			}
		} else {
			dfs(y, x)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}

	p := 0
	d := 0
	c := 0
	for i := 1; i <= n; i++ {
		if len(e[i]) == 0 {
			p++
		} else if k[i] == 0 {
			f = 0
			dfs(i, f)
			if f != 0 {
				d++
			} else {
				c++
			}
		}
	}
	fmt.Println(2*n*p - p*p + c*c*2 + d*d + c*d*2)
}
