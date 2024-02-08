package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 10005
const M = 1005
const P = 998244353

var lc, rc, fa, dep [N]int
var n, m, K int
var ffa [N][15]int

func dfs1(u int) {
	ffa[u][0] = fa[u]
	dep[u] = dep[fa[u]] + 1
	for i := 1; i <= 13; i++ {
		ffa[u][i] = ffa[ffa[u][i-1]][i-1]
	}
	if lc[u] != 0 {
		dfs1(lc[u])
	}
	if rc[u] != 0 {
		dfs1(rc[u])
	}
}

func LCA(u, v int) int {
	if dep[u] < dep[v] {
		u, v = v, u
	}
	kd := dep[u] - dep[v]
	for i := 13; i >= 0; i-- {
		if ((kd >> i) & 1) != 0 {
			u = ffa[u][i]
		}
	}
	if u == v {
		return u
	}
	for i := 13; i >= 0; i-- {
		if ffa[u][i] != ffa[v][i] {
			u = ffa[u][i]
			v = ffa[v][i]
		}
	}
	return fa[u]
}

func K_fa(u, kd int) int {
	for i := 13; i >= 0; i-- {
		if ((kd >> i) & 1) != 0 {
			u = ffa[u][i]
		}
	}
	return u
}

var dt, dt2, md, md2, fx, sz [N]int
var f [N][M]int
var bf [M]int

func dfs2(u int) {
	if lc[u] != 0 {
		dfs2(lc[u])
	}
	if rc[u] != 0 {
		dfs2(rc[u])
	}
	sz[u] = md[u] + sz[lc[u]] + sz[rc[u]] - dt[u] - dt2[u] + md2[u]
	for i := 0; i <= K; i++ {
		f[u][i] = 1
	}
	for i := 0; i <= K; i++ {
		if i+md[u]-dt[u] < 0 || i-dt[u]+md[u] > K || i+md[u] > K {
			f[u][i] = 0
		}
	}
	if lc[u] != 0 {
		if rc[u] == 0 {
			for i := 0; i <= K; i++ {
				if f[u][i] != 0 {
					f[u][i] = f[lc[u]][i-dt[u]+md[u]]
				}
			}
		} else {
			for i := 0; i <= K; i++ {
				bf[i] = f[u][i]
				f[u][i] = 0
			}
			if fx[u] != 1 {
				for i := 0; i <= K; i++ {
					if bf[i] != 0 {
						bx := i + md[u] - dt[u] + sz[lc[u]]
						if bx >= 0 && bx <= K {
							f[u][i] = (f[u][i] + f[lc[u]][i+md[u]-dt[u]]*f[rc[u]][bx]%P) % P
						}
					}
				}
			}
			if fx[u] != 0 {
				for i := 0; i <= K; i++ {
					if bf[i] != 0 {
						bx := i + md[u] - dt[u] + sz[rc[u]]
						if bx >= 0 && bx <= K {
							f[u][i] = (f[u][i] + f[rc[u]][i+md[u]-dt[u]]*f[lc[u]][bx]%P) % P
						}
					}
				}
			}
		}
	}
	bx := md[u] + sz[lc[u]] + sz[rc[u]] - dt[u] + md2[u]
	for i := 0; i <= K; i++ {
		if i+sz[u] < 0 || i+sz[u] > K || i+bx < 0 || i+bx > K {
			f[u][i] = 0
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &K)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &fa[i])
		if lc[fa[i]] != 0 {
			rc[fa[i]] = i
		} else {
			lc[fa[i]] = i
		}
	}
	for i := 1; i <= n; i++ {
		fx[i] = 2
	}
	dfs1(1)
	for m > 0 {
		m--
		var u, v int
		fmt.Fscan(in, &u, &v)
		x := LCA(u, v)
		if x == u {
			y := K_fa(v, dep[v]-dep[u]-1)
			if y == rc[u] {
				md[rc[u]]++
			} else {
				md[lc[u]]++
			}
			dt[v]++
		} else if x == v {
			md2[u]++
			y := K_fa(u, dep[u]-dep[v]-1)
			dt2[y]++
		} else {
			yu := K_fa(u, dep[u]-dep[x]-1)
			bx := 0
			if yu == lc[x] {
				bx = 0
			} else {
				bx = 1
			}
			if fx[x] != 2 && fx[x] != bx {
				fmt.Println(0)
				return
			}
			fx[x] = bx
			md2[u]++
			dt[v]++
		}
	}
	dfs2(1)
	fmt.Println(f[1][0])
}
