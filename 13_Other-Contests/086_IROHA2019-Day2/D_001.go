package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var par, sz [200005]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type tuple struct {
		x, y, z, a int
	}

	var N, M int
	fmt.Fscan(in, &N, &M)

	for i := 0; i < N+1; i++ {
		par[i] = i
		sz[i] = 1
	}

	edge := make([]tuple, 0)
	for i := 1; i <= M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		edge = append(edge, tuple{-c, a, b, i})
	}
	sort.Slice(edge, func(i, j int) bool {
		if edge[i].x == edge[j].x {
			if edge[i].y == edge[j].y {
				return edge[i].z < edge[j].z
			}
			return edge[i].y < edge[j].y
		}
		return edge[i].x < edge[j].x
	})

	ans := make([]int, 0)
	for _, e := range edge {
		_, a, b, i := e.x, e.y, e.z, e.a
		a = find(a)
		b = find(b)
		if a != b {
			join(a, b)
			ans = append(ans, i)
		}
	}
	sort.Ints(ans)
	for _, i := range ans {
		fmt.Fprintln(out, i)
	}
}

func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

func join(x, y int) {
	if sz[x] < sz[y] {
		x, y = y, x
	}
	sz[x] += sz[y]
	par[y] = x
}
