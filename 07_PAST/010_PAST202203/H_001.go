package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var G [200200][]int

	var n, q int
	fmt.Fscan(in, &n, &q)
	par := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
	}
	for i := 0; i < n; i++ {
		G[i] = append(G[i], i)
	}
	for q > 0 {
		q--
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var u, v int
			fmt.Fscan(in, &u, &v)
			u--
			v--
			if par[u] == par[v] {
				continue
			}
			if len(G[par[u]]) < len(G[par[v]]) {
				u, v = v, u
			}
			for _, nv := range G[par[v]] {
				G[par[u]] = append(G[par[u]], nv)
				par[nv] = par[u]
			}
		} else {
			var u int
			fmt.Fscan(in, &u)
			u--
			sort.Ints(G[par[u]])
			k := len(G[par[u]])
			for i := 0; i < k; i++ {
				if i != 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, G[par[u]][i]+1)
			}
			fmt.Fprintln(out)
		}
	}
}
