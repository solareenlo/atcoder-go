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

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	G := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	C := make([]int, n)
	for i := range C {
		fmt.Fscan(in, &C[i])
	}

	for i := 0; i < q; i++ {
		var t, x int
		fmt.Fscan(in, &t, &x)
		x--
		if t == 1 {
			fmt.Fprintln(out, C[x])
			for _, v := range G[x] {
				C[v] = C[x]
			}
		} else {
			var color int
			fmt.Fscan(in, &color)
			fmt.Fprintln(out, C[x])
			C[x] = color
		}
	}
}
