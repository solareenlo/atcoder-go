package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 600010
const INF = 1000000000

var n, m int
var cnt, cnt2 int
var vt, nt [N][]int
var vis [N]bool
var stk []int
var mxd, dist, dfn, lw, dp, sz [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		fmt.Fscan(in, &n, &m)
		for i := 0; i < n; i++ {
			vt[i] = make([]int, 0)
		}
		for i := 0; i < m; i++ {
			var x, y int
			fmt.Fscan(in, &x, &y)
			x--
			y--
			vt[x] = append(vt[x], y)
			vt[y] = append(vt[y], x)
		}
		for i := 0; i < cnt2+6; i++ {
			nt[i] = make([]int, 0)
			vis[i] = false
		}
		cnt = 0
		cnt2 = n
		for len(stk) != 0 {
			stk = stk[:len(stk)-1]
		}
		predfs(0, -1)
		mxl := Solve()
		fmt.Println((max(CALC(mxl.x), CALC(mxl.y)) + n) / 2)
	}
}

func predfs(x, lst int) {
	tot := 0
	stk = append(stk, x)
	vis[x] = true
	cnt++
	dfn[x] = cnt
	lw[x] = cnt
	for i := 0; i < len(vt[x]); i++ {
		if !vis[vt[x][i]] {
			tot++
			predfs(vt[x][i], x)
			lw[x] = min(lw[x], lw[vt[x][i]])
			if lw[vt[x][i]] >= dfn[x] {
				for stk[len(stk)-1] != vt[x][i] {
					nt[cnt2] = append(nt[cnt2], stk[len(stk)-1])
					nt[stk[len(stk)-1]] = append(nt[stk[len(stk)-1]], cnt2)
					stk = stk[:len(stk)-1]
				}
				nt[cnt2] = append(nt[cnt2], stk[len(stk)-1])
				nt[stk[len(stk)-1]] = append(nt[stk[len(stk)-1]], cnt2)
				stk = stk[:len(stk)-1]
				nt[cnt2] = append(nt[cnt2], x)
				nt[x] = append(nt[x], cnt2)
				cnt2++
			}
		} else if vt[x][i] != lst {
			lw[x] = min(lw[x], dfn[vt[x][i]])
		}
	}
	if lst == -1 && tot == 0 {
		nt[cnt2] = append(nt[cnt2], x)
		nt[x] = append(nt[x], cnt2)
		cnt2++
	}
	return
}

type pair struct {
	x, y int
}

func Solve() pair {
	x := mxdist(0)
	y := mxdist(x)
	return pair{x, y}
}

func mxdist(st int) int {
	q := make([]int, 0)
	for i := 0; i < cnt2; i++ {
		vis[i] = false
	}
	q = append(q, st)
	vis[st] = true
	ret := st
	for len(q) != 0 {
		x := q[0]
		q = q[1:]
		for i := 0; i < len(nt[x]); i++ {
			if !vis[nt[x][i]] {
				vis[nt[x][i]] = true
				ret = nt[x][i]
				q = append(q, nt[x][i])
			}
		}
	}
	return ret
}

func CALC(st int) int {
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		vis[i] = false
		dist[i] = 0
	}
	q = append(q, st)
	vis[st] = true
	dist[st] = 0
	for len(q) != 0 {
		x := q[0]
		q = q[1:]
		for i := 0; i < len(vt[x]); i++ {
			if !vis[vt[x][i]] {
				vis[vt[x][i]] = true
				dist[vt[x][i]] = dist[x] + 1
				q = append(q, vt[x][i])
			}
		}
	}
	return dfs(st, -1)
}

func dfs(x, lst int) int {
	ret := 0
	dp[x] = 0
	if x < n {
		sz[x] = 1
	} else {
		sz[x] = 0
	}
	tmp := 0
	if lst != -1 {
		tmp = 1
	}
	if len(nt[x])-tmp == 0 {
		return 0
	}
	for i := 0; i < len(nt[x]); i++ {
		if nt[x][i] != lst {
			ret = max(ret, dfs(nt[x][i], x))
			sz[x] += sz[nt[x][i]]
		}
	}
	if x < n {
		dp[x] = -INF
		for i := 0; i < len(nt[x]); i++ {
			y := nt[x][i]
			if y != lst {
				for j := 0; j < len(nt[y]); j++ {
					if nt[y][j] != x {
						ret = max(ret, dp[nt[y][j]]+(dist[nt[y][j]]-dist[x])*2+sz[nt[y][j]]-sz[y]-1)
						dp[x] = max(dp[x], dp[nt[y][j]]+(dist[nt[y][j]]-dist[x])*2+sz[nt[y][j]]-sz[x])
					}
				}
			}
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
