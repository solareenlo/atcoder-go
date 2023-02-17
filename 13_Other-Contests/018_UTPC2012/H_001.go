package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1 << 30
const MAX_N = 100010
const MAX_E = 20

type P struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, query int
	fmt.Fscan(in, &n, &query)
	p := make([]P, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})
	p = append(p, P{INF, INF})
	tails := make([]int, n+1)
	tails[n] = n
	for i := n - 1; i >= 0; i-- {
		tails[i] = tails[i+1]
		if p[tails[i]].y > p[i].y {
			tails[i] = i
		}
	}
	next := make([][]int, MAX_E)
	for i := range next {
		next[i] = make([]int, MAX_N)
	}
	for i := 0; i < n; i++ {
		next[0][i] = tails[lowerBound(p, P{p[i].y, -INF})]
	}
	next[0][n] = n
	for e := 0; e < 18; e++ {
		for i := 0; i <= n; i++ {
			next[e+1][i] = next[e][next[e][i]]
		}
	}
	for query > 0 {
		query--
		var a, b int
		fmt.Fscan(in, &a, &b)
		i := tails[lowerBound(p, P{a, -INF})]
		if p[i].y > b {
			fmt.Println(0)
			continue
		}
		res := 0
		for e := 17; e >= 0; e-- {
			if p[next[e][i]].y <= b {
				res |= 1 << e
				i = next[e][i]
			}
		}
		fmt.Println(res + 1)
	}
}

func lowerBound(a []P, x P) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}
