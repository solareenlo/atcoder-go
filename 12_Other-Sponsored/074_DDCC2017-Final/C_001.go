package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1 << 60)

type edge struct {
	to, cost int
}

var N, M int
var g [300][]edge
var v []int

func dfs(idx, par int) int {
	ret := 0
	for _, e := range g[idx] {
		if v[e.to] == INF {
			v[e.to] = v[idx] + e.cost
			ret += dfs(e.to, par)
		} else if e.to != par {
			if v[idx]+e.cost != v[e.to] {
				return (114514)
			}
		} else {
			if v[idx]+e.cost != v[e.to] {
				ret++
			}
		}
	}
	return (ret)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		g[a] = append(g[a], edge{b, c})
	}

	for i := 0; i < N; i++ {
		v = make([]int, N)
		for j := range v {
			v[j] = INF
		}
		v[i] = 0
		if dfs(i, i) <= 1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
