package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)
	g := make([][]NODE, N)
	for i := 0; i < N-1; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		g[a] = append(g[a], NODE{a, b, c})
		g[b] = append(g[b], NODE{b, a, c})
	}
	l, r := 0, 1000000010
	for r-l > 1 {
		m := (l + r) / 2
		if DFS(g, 0, -1, m).cnt >= K {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
}

type Info struct {
	dist, cnt int
}

type NODE struct {
	src, dst, l int
}

func DFS(g [][]NODE, pos, prev, d int) Info {
	var res Info
	near, far := 0, 1000000010
	for _, n := range g[pos] {
		if n.dst == prev {
			continue
		}
		ni := DFS(g, n.dst, pos, d)
		res.cnt += ni.cnt
		nd := ni.dist + n.l
		if 2*nd < d {
			near = max(near, nd)
			res.cnt--
		} else {
			far = min(far, nd)
		}
	}
	if near+far < d {
		res.dist = far
	} else {
		res.dist = near
		res.cnt++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
