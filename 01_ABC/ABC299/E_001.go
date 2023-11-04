package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 2020

	var e [N][]int
	var c, q, a [N]int
	var f [N][N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}
	var k int
	fmt.Fscan(in, &k)
	for i := range c {
		c[i] = -1
	}
	for i := 1; i <= k; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		c[u] = v
	}
	for p := 1; p <= n; p++ {
		l, r := 1, 1
		f[p][p] = 1
		q[1] = p
		for l <= r {
			x := q[l]
			l++
			if f[p][x] <= c[p] {
				a[x] = 1
			}
			for _, i := range e[x] {
				if f[p][i] == 0 {
					f[p][i] = f[p][x] + 1
					r++
					q[r] = i
				}
			}
		}
	}
	ans := true
	for p := 1; p <= n; p++ {
		g := c[p] < 0
		for i := 1; i <= n; i++ {
			if !g && !(f[p][i] == c[p]+1 && a[i] == 0) {
				g = false
			} else {
				g = true
			}
		}
		if ans && g {
			ans = true
		} else {
			ans = false
		}
	}
	if ans {
		fmt.Println("Yes")
		for i := 1; i <= n; i++ {
			fmt.Printf("%d", a[i]^1)
		}
	} else {
		fmt.Println("No")
	}
}
