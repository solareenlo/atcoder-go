package main

import (
	"fmt"
	"sort"
)

const INF = 1000000000

type Edge struct {
	fr, sc int
}

type Pair struct {
	first  int
	second Edge
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	es := make([]Pair, m)
	largest := make([]int, 45)
	val := make([]int, 45)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &es[i].second.fr, &es[i].second.sc, &es[i].first)
		es[i].second.fr--
		es[i].second.sc--
		largest[es[i].second.fr] = max(largest[es[i].second.fr], es[i].first)
		largest[es[i].second.sc] = max(largest[es[i].second.sc], es[i].first)
	}

	sort.Slice(es, func(i, j int) bool {
		return es[i].first < es[j].first
	})

	for i, j := 0, m-1; i < j; {
		es[i], es[j] = es[j], es[i]
		i++
		j--
	}

	root := make([]int, 0)
	for i := 0; i < m; i++ {
		suc := true
		a, b := es[i].second.fr, es[i].second.sc
		for j := 0; j < i; j++ {
			if es[j].second.fr == a || es[j].second.sc == a || es[j].second.fr == b || es[j].second.sc == b {
				suc = false
			}
		}
		if suc {
			root = append(root, a)
		}
	}

	k := len(root)
	res := 0
	ans := make([]int, 45)
	for bit := 0; bit < (1 << k); bit++ {
		for i := 0; i < n; i++ {
			val[i] = INF
		}
		for i := 0; i < k; i++ {
			if bit>>i&1 == 1 {
				val[root[i]] = largest[root[i]]
			}
		}
		cnt := 0
		for i := 0; i < m; i++ {
			a, b, c := es[i].second.fr, es[i].second.sc, es[i].first
			if (val[a] == INF) != (val[b] == INF) {
				if val[a] == INF {
					val[a] = c
				}
				if val[b] == INF {
					val[b] = c
				}
				cnt++
			}
		}
		if res < cnt {
			res = cnt
			copy(ans, val)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
