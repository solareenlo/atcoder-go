package main

import "fmt"

const N = 200001

var (
	g [N][26]int
	f [N]int
	l [N]int
	c [N]int
	n int
)

func split(p, o int) int {
	q := g[p][o]
	if l[p]+1 == l[q] {
		return q
	}
	n++
	nq := n
	l[nq] = l[p] + 1
	copy(g[nq][:], g[q][:])
	f[nq] = f[q]
	f[q] = nq
	for ; p != 0 && g[p][o] == q; p = f[p] {
		g[p][o] = nq
	}
	return nq
}

func extend(r, p, o int) int {
	n++
	np := n
	l[np] = l[p] + 1
	for i := range g[np] {
		g[np][i] = 0
	}
	for ; p != 0 && g[p][o] == 0; p = f[p] {
		g[p][o] = np
	}
	f[np] = r
	if p != 0 {
		f[np] = split(p, o)
	}
	return np
}

func main() {
	var s string
	fmt.Scan(&s)
	r := 1
	n = 1
	p := r
	for _, ch := range s {
		p = extend(r, p, int(ch-'a'))
	}
	for i := 2; i <= n; i++ {
		c[l[f[i]]+1]++
		c[l[i]+1]--
	}
	for i := 1; i < len(s); i++ {
		c[i+1] += c[i]
	}
	ans := 0
	for i := 1; i <= len(s); i++ {
		ans = max(ans, c[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
