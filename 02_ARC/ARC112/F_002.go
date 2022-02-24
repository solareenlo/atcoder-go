package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n int
	c = [16]int{}
)

func trans() int {
	r := 0
	for i := n - 1; i >= 0; i-- {
		r = r * 2 * (i + 1)
		r += c[i]
	}
	return r
}

var (
	ans int = 10000
	set     = [16]int{}
)

func dfs(d, r, s, u, v, m int) {
	if ans <= r {
		return
	}
	if d == -1 {
		if r != 0 {
			ans = r
		}
		return
	}
	m /= 2 * (d + 1)
	for i := 0; i < 2*(d+1); i++ {
		p := (s+m-1+v-u)/v - 1
		if s <= p*v+u {
			set[d] = i
			dfs(d-1, r+i, s, u, v, m)
		}
		s += m
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &n, &m)

	mod := 1
	for i := 1; i <= n; i++ {
		mod *= 2 * i
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}
	p := trans() % mod

	g := mod - 1
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &c[i])
		}
		g = gcd(g, trans())
	}
	p %= g

	dfs(n-1, 0, 0, p, g, mod)
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
