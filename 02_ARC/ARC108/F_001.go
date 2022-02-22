package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 400004
const mod = 1_000_000_007

var (
	sze int
	to  = make([]int, N)
	nx  = make([]int, N)
	hd  = make([]int, N)
)

func add(u, v int) {
	sze++
	to[sze] = v
	nx[sze] = hd[u]
	hd[u] = sze
}

func DFS(u, p int, d []int) int {
	mx := u
	for i := hd[u]; i > 0; i = nx[i] {
		v := to[i]
		if v != p {
			d[v] = d[u] + 1
			w := DFS(v, u, d)
			if d[w] > d[mx] {
				mx = w
			}
		}
	}
	return mx
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		add(u, v)
		add(v, u)
	}

	S := make([]int, N)
	s := DFS(1, 0, S)
	S[s] = 0
	t := DFS(s, 0, S)
	T := make([]int, N)
	DFS(t, 0, T)
	u := 0
	c := make([]int, N)
	for i := 1; i <= n; i++ {
		u = max(u, min(S[i], T[i]))
		c[max(S[i], T[i])]++
	}

	p := make([]int, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		c[i] += c[i-1]
		p[i] = 2 * p[i-1] % mod
	}

	ans := 0
	v := 0
	for i := T[s]; i > 0; i-- {
		var w int
		if i <= u {
			w = p[n]
		} else {
			w = (p[n] - 2*p[c[i-1]]%mod + mod) % mod
		}
		ans = (ans + i*(w-v+mod)) % mod
		v = w
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
