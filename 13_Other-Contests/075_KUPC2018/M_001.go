package main

import (
	"bufio"
	"fmt"
	"os"
)

const NN = 140000
const S = 3000

type tuple struct {
	x, y, z int
}

var n, q, k int
var g [NN][]int
var ADD, qry [NN + 1][]tuple
var Ans [NN]int
var sz, head, vid, par, L, R, et [NN]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	dfs(0, -1)
	dfs2(0, -1, 0)
	fmt.Fscan(in, &q)
	root := 0
	for i := 0; i < q; i++ {
		Ans[i] = -1
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var v, x, k int
			fmt.Fscan(in, &v, &x, &k)
			v--
			if v == root {
				ADD[0] = append(ADD[0], tuple{x, k, i})
				ADD[n] = append(ADD[n], tuple{x, -k, i})
			} else if lca(v, root) != v {
				ADD[L[v]] = append(ADD[L[v]], tuple{x, k, i})
				ADD[R[v]] = append(ADD[R[v]], tuple{x, -k, i})
			} else {
				u := lift(v, root)
				ADD[0] = append(ADD[0], tuple{x, k, i})
				ADD[n] = append(ADD[n], tuple{x, -k, i})
				ADD[L[u]] = append(ADD[L[u]], tuple{x, -k, i})
				ADD[R[u]] = append(ADD[R[u]], tuple{x, k, i})
			}
		} else if t == 2 {
			var v, y, z int
			fmt.Fscan(in, &v, &y, &z)
			v--
			qry[L[v]] = append(qry[L[v]], tuple{y, z, i})
		} else {
			var v int
			fmt.Fscan(in, &v)
			v--
			root = v
		}
	}
	for l := 0; l < q; l += S {
		each := make([]tuple, 0)
		r := min(l+S, q)
		for i := 0; i <= n; i++ {
			for _, p := range ADD[i] {
				x, k, t := p.x, p.y, p.z
				if t < l {
					insert(x, k)
				} else if t < r {
					each = append(each, p)
				}
			}
			for _, p := range qry[i] {
				y, z, t := p.x, p.y, p.z
				if l <= t && t < r {
					Ans[t] = query(y, z)
					for _, e := range each {
						if e.z < t {
							if (e.x ^ y) <= z {
								Ans[t] += e.y
							}
						}
					}
				}
			}
		}
	}
	for i := 0; i < q; i++ {
		if Ans[i] != -1 {
			fmt.Println(Ans[i])
		}
	}
}

func dfs(u, p int) {
	sz[u] = 1
	for _, v := range g[u] {
		if v != p {
			dfs(v, u)
			sz[u] += sz[v]
		}
	}
}

func dfs2(u, p, h int) {
	et[k] = u
	L[u] = k
	vid[u] = k
	k++
	head[u] = h
	par[u] = p
	for _, v := range g[u] {
		if v != p && sz[u] < sz[v]*2 {
			dfs2(v, u, h)
		}
	}
	for _, v := range g[u] {
		if v != p && sz[u] >= sz[v]*2 {
			dfs2(v, u, v)
		}
	}
	R[u] = k
}

func lca(u, v int) int {
	for head[u] != head[v] {
		if vid[u] < vid[v] {
			u, v = v, u
		}
		u = par[head[u]]
	}
	if vid[u] < vid[v] {
		return u
	}
	return v
}

func lift(u, v int) int {
	for {
		if head[u] == head[v] {
			return et[vid[u]+1]
		}
		if u == par[head[v]] {
			return head[v]
		}
		v = par[head[v]]
	}
}

var succ [30*280000 + 1][2]int
var ptr int
var cnt [30*280000 + 1]int

func insert(k, v int) {
	x := 0
	for i := 29; i >= 0; i-- {
		y := (k >> i) & 1
		if succ[x][y] == 0 {
			ptr++
			succ[x][y] = ptr
		}
		x = succ[x][y]
		cnt[x] += v
	}
}

func query(y, z int) int {
	res := 0
	x := 0
	for i := 29; i >= 0; i-- {
		if ((^y >> i) & 1) != 0 {
			if ((z >> i) & 1) != 0 {
				res += cnt[succ[x][0]]
				x = succ[x][1]
			} else {
				x = succ[x][0]
			}
		} else {
			if ((z >> i) & 1) != 0 {
				res += cnt[succ[x][1]]
				x = succ[x][0]
			} else {
				x = succ[x][1]
			}
		}
		if x == 0 {
			break
		}
	}
	res += cnt[x]
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
