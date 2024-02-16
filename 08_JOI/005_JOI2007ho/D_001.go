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

	const N = 5000

	var n, m int
	fmt.Fscan(in, &n, &m)
	G := make([][]int, N)
	var C [N]int
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		C[b]++
	}
	Q := make([]int, 0)
	for i := 0; i < n; i++ {
		if C[i] == 0 {
			Q = append(Q, i)
		}
	}
	f := 0
	for len(Q) != 0 {
		if len(Q) > 1 {
			f = 1
		}
		v := Q[0]
		Q = Q[1:]
		fmt.Fprintln(out, v+1)
		for _, u := range G[v] {
			C[u]--
			if C[u] == 0 {
				Q = append(Q, u)
			}
		}
	}
	fmt.Fprintln(out, f)
}
