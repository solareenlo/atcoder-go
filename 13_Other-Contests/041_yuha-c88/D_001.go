package main

import (
	"fmt"
)

type pair struct {
	first, second int
}

var (
	S, G, mini int = 0, 0, 2e9
	g              = make([][]pair, 9)
	used           = make([]int, 32)
	s              = make([]string, 9)
	a, res     string
)

func dfs(cur, dir, dist int) {
	if cur == G {
		dir = 1
	}
	if cur == S && dir == 1 {
		if dist < mini || a < res {
			mini = dist
			res = a
		}
		return
	}
	if dist >= mini {
		return
	}
	for _, i := range g[cur] {
		if used[i.second] == 0 {
			used[i.second] = 1
			a += s[i.first]
			dfs(i.first, dir, dist+1)
			a = a[:len(a)-len(s[i.first])]
			used[i.second] = 0
		}
	}
}

func main() {
	var n, M int
	fmt.Scan(&n)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
		m[s[i]] = i
	}
	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		var a, b string
		fmt.Scan(&a, &b)
		x, y := m[a], m[b]
		g[x] = append(g[x], pair{y, i})
		g[y] = append(g[y], pair{x, i})
	}
	var c, d string
	fmt.Scan(&c, &d)
	S, G = m[c], m[d]
	a = s[S]
	dfs(S, 0, 0)
	res = res[:len(res)-len(s[S])]
	fmt.Println(res)
}
