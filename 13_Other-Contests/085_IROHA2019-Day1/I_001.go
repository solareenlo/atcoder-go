package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 100005
	const inf = int(1e9)

	type T struct {
		a, c, e int
	}

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	G := make([][]T, MX)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], T{b, c, 0})
		G[b] = append(G[b], T{a, c, 0})
	}
	b := MX
	var d [MX]int
	for i := 1; i < n; i++ {
		d[i] = inf
	}
	var qu [MX * 2]T
	e := MX + 1
	for b < e {
		t := qu[b]
		b++
		if d[t.a] < t.e {
			continue
		}
		for _, nx := range G[t.a] {
			tmp := 0
			if nx.c != t.c {
				tmp = 1
			}
			if (t.e + tmp) < d[nx.a] {
				nx.e = t.e + tmp
				d[nx.a] = nx.e
				if nx.c == t.c {
					b--
					qu[b] = nx
				} else {
					qu[e] = nx
					e++
				}
			}
		}
	}
	if d[n-1] == inf {
		fmt.Println(-1)
	} else {
		fmt.Println(d[n-1] * k)
	}
}
