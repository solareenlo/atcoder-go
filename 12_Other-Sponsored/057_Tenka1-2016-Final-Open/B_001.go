package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type graph [][]int

func add_undirected_edge(G *[][]int, u, v int) {
	(*G)[u] = append((*G)[u], v)
	(*G)[v] = append((*G)[v], u)
}

func connected_components(G [][]int) [][]int {
	n := len(G)
	res := make([][]int, 0)
	vis := make([]bool, n)
	for u := 0; u < n; u++ {
		if !vis[u] {
			vis[u] = true
			res = append(res, make([]int, 0))
			Q := make([]int, 0)
			Q = append(Q, u)
			for len(Q) > 0 {
				v := Q[0]
				Q = Q[1:]
				res[len(res)-1] = append(res[len(res)-1], v)
				for _, w := range G[v] {
					if !vis[w] {
						vis[w] = true
						Q = append(Q, w)
					}
				}
			}
		}
	}
	return res
}

func solve(G [][]int) []int {
	n := len(G)
	order := make([]int, n)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		H := make([][]int, n)
		for u := 0; u < n; u++ {
			for _, v := range G[u] {
				if !used[u] && !used[v] && u < v {
					add_undirected_edge(&H, u, v)
				}
			}
		}
		u := -1
		for v := 0; v < n; v++ {
			if !used[v] {
				if u == -1 || len(H[u]) > len(H[v]) {
					u = v
				}
			}
		}
		order[n-i-1] = u
		used[u] = true
	}

	color := make([]int, n)
	for i := range color {
		color[i] = -1
	}

	type pair struct {
		x, y int
	}

	var Kempe func(int) bool
	Kempe = func(u int) bool {
		for i := 0; i < 100; i++ {
			nbd := make([]int, 0)
			for _, v := range G[u] {
				if color[v] != -1 {
					nbd = append(nbd, v)
				}
			}
			rand.Shuffle(len(nbd), func(i, j int) { nbd[i], nbd[j] = nbd[j], nbd[i] })
			for _, v := range nbd {
				for c1 := 0; c1 < 4; c1++ {
					c0 := color[v]
					if c1 != c0 {
						S := make([]pair, 1)
						S[0] = pair{v, c0}
						color[v] = -2
						Q := make([]int, 0)
						Q = append(Q, v)
						for len(Q) > 0 {
							w := Q[0]
							Q = Q[1:]
							for _, z := range G[w] {
								if color[z] == c0 || color[z] == c1 {
									S = append(S, pair{z, color[z]})
									color[z] = -2
									Q = append(Q, z)
								}
							}
						}
						for _, tmp := range S {
							w := tmp.x
							c := tmp.y
							if c == c0 {
								color[w] = c1
							} else {
								color[w] = c0
							}
						}
					}
					var used [4]bool
					for _, w := range nbd {
						used[color[w]] = true
					}
					for c2 := 0; c2 < 4; c2++ {
						if !used[c2] {
							color[u] = c2
							return true
						}
					}
				}
			}
		}
		return false
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == n {
			return true
		}
		u := order[i]
		var used [4]bool
		for _, v := range G[u] {
			if color[v] != -1 {
				used[color[v]] = true
			}
		}
		for c := 0; c < 4; c++ {
			if !used[c] {
				color[u] = c
				if dfs(i + 1) {
					return true
				}
				color[u] = -1
			}
		}
		if Kempe(u) {
			if dfs(i + 1) {
				return true
			}
			color[u] = -1
		}
		return false
	}
	dfs(0)
	return color
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	G := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		add_undirected_edge(&G, u, v)
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for _, C := range connected_components(G) {
		N := len(C)
		f := make(map[int]int)
		for i := 0; i < N; i++ {
			f[C[i]] = i
		}

		H := make([][]int, N)
		for _, u := range C {
			for _, v := range G[u] {
				i := f[u]
				j := f[v]
				if i < j {
					add_undirected_edge(&H, i, j)
				}
			}
		}

		color := solve(H)
		for i := 0; i < N; i++ {
			ans[C[i]] = color[i]
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(ans[i] + 1)
	}
}
