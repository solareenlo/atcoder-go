package main

import (
	"fmt"
	"math/bits"
)

var (
	u, v    int
	G       = [55][]edge{}
	B       = [21]int{}
	visited = [55]bool{}
)

type edge struct{ to, id int }

func dfs(i, k int) bool {
	visited[k] = false
	if k == v {
		return true
	}
	for _, e := range G[k] {
		if visited[e.to] && dfs(i, e.to) {
			B[i] |= 1 << e.id
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		G[a] = append(G[a], edge{b, i})
		G[b] = append(G[b], edge{a, i})
	}

	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&u, &v)
		u--
		v--
		for j := range visited {
			visited[j] = true
		}
		dfs(i, u)
	}

	res := 0
	for i := 1<<m - 1; i >= 0; i-- {
		a := 0
		for j := m; j >= 0; j-- {
			a |= B[j] * (i >> j & 1)
		}
		b := 1 << (n - 1 - bits.OnesCount(uint(a)))
		if bits.OnesCount(uint(i))&1 != 0 {
			res -= b
		} else {
			res += b
		}
	}
	fmt.Println(res)
}
