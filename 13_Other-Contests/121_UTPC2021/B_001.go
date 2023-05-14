package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

type P struct {
	x pair
	y int
}

var G [200010][]pair
var col [200010]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	var a, b [200010]int
	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}
		for i := 0; i < n; i++ {
			G[i] = make([]pair, 0)
			if a[i] > b[i] {
				a[i], b[i] = b[i], a[i]
			}
		}
		v := make([]P, 0)
		for i := 0; i < n; i++ {
			v = append(v, P{pair{a[i], 1}, i})
			v = append(v, P{pair{b[i], -1}, i})
		}
		sortP(v)
		for i := 0; i < n; i++ {
			col[i] = -1
		}
		for i := 0; i < 2*n; i += 2 {
			if v[i].y == v[i+1].y {
				continue
			}
			if v[i+1].x.y == v[i].x.y {
				G[v[i+1].y] = append(G[v[i+1].y], pair{v[i].y, 1})
				G[v[i].y] = append(G[v[i].y], pair{v[i+1].y, 1})
			} else {
				G[v[i+1].y] = append(G[v[i+1].y], pair{v[i].y, 0})
				G[v[i].y] = append(G[v[i].y], pair{v[i+1].y, 0})
			}
		}
		for i := 0; i < n; i++ {
			if col[i] == -1 {
				col[i] = 0
				dfs(i, -1)
			}
		}
		for i := 0; i < n; i++ {
			if col[i] == 1 {
				fmt.Fprintf(out, "%d ", a[i])
			} else {
				fmt.Fprintf(out, "%d ", b[i])
			}
		}
		fmt.Fprintln(out)
		for i := 0; i < n; i++ {
			if col[i] == 1 {
				fmt.Fprintf(out, "%d ", b[i])
			} else {
				fmt.Fprintf(out, "%d ", a[i])
			}
		}
		fmt.Fprintln(out)
	}
}

func dfs(s, p int) {
	for _, e := range G[s] {
		if col[e.x] != -1 {
			continue
		}
		col[e.x] = col[s] ^ e.y
		dfs(e.x, s)
	}
}

func sortP(tup []P) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].x.x == tup[j].x.x {
			if tup[i].x.y == tup[j].x.y {
				return tup[i].y < tup[j].y
			}
			return tup[i].x.y < tup[j].x.y
		}
		return tup[i].x.x < tup[j].x.x
	})
}
