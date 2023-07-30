package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	t, c int
}

var G [1001][]edge

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, s, t int
	fmt.Fscan(in, &n, &m, &s, &t)
	for i := 0; i < m; i++ {
		var x, y, d int
		fmt.Fscan(in, &x, &y, &d)
		G[x] = append(G[x], edge{y, d})
		G[y] = append(G[y], edge{x, d})
	}
	a := bf(s)
	b := bf(t)
	for i := 1; i <= n; i++ {
		if a[i] == b[i] && a[i] < 1<<44 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}

func bf(s int) []int {
	d := make([]int, 1111)
	for i := range d {
		d[i] = 1 << 55
	}
	d[s] = 0
	for {
		f := false
		for i := 0; i < 1001; i++ {
			for _, e := range G[i] {
				if d[e.t] > d[i]+e.c {
					f = true
					d[e.t] = d[i] + e.c
				}
			}
		}
		if !f {
			break
		}
	}
	return d
}
