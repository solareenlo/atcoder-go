package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 200000
	const B = 632

	var n, m int
	fmt.Fscan(in, &n, &m)
	var G [N][]int
	for m > 0 {
		m--
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	var big [N]bool
	for i := 0; i < n; i++ {
		big[i] = len(G[i]) > B
	}

	var G2 [N][]int
	for i := 0; i < n; i++ {
		for _, u := range G[i] {
			if big[u] {
				G2[i] = append(G2[i], u)
			}
		}
	}

	var q int
	fmt.Fscan(in, &q)

	var ys [N + 1]int
	var t, col [N]int
	ys[q] = 1
	for i := 0; i < n; i++ {
		t[i] = q
		col[i] = 1
	}

	for q > 0 {
		q--
		var v, y int
		fmt.Fscan(in, &v, &y)
		v--
		if !big[v] {
			mini := t[v]
			for _, u := range G[v] {
				mini = min(mini, t[u])
			}
			col[v] = ys[mini]
		}
		fmt.Fprintln(out, col[v])
		ys[q] = y
		t[v] = q
		col[v] = y
		for _, u := range G2[v] {
			col[u] = y
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
