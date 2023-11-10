package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200020

type edge struct {
	to, nxt int
}

var cnt, tim, tot int
var e [N << 1]edge
var head, low, dfn, siz, cut [N]int
var stk []int
var qwq [N][]int

func addedge(u, v int) {
	cnt++
	e[cnt].to = v
	e[cnt].nxt = head[u]
	head[u] = cnt
}

func Tarjan(u, fa int) {
	tim++
	low[u] = tim
	dfn[u] = low[u]
	stk = append(stk, u)
	for i := head[u]; i > 0; i = e[i].nxt {
		v := e[i].to
		if v == fa {
			continue
		}
		if dfn[v] == 0 {
			Tarjan(v, u)
			low[u] = min(low[u], low[v])
			if dfn[u] <= low[v] {
				tot++
				siz[tot] += 2
				cut[u]++
				cut[v]++
				qwq[tot] = append(qwq[tot], u)
				qwq[tot] = append(qwq[tot], v)
				for stk[len(stk)-1] != v {
					siz[tot]++
					cut[stk[len(stk)-1]]++
					qwq[tot] = append(qwq[tot], stk[len(stk)-1])
					stk = stk[:len(stk)-1]
				}
				stk = stk[:len(stk)-1]
			}
		} else {
			low[u] = min(low[u], dfn[v])
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var deg [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		deg[x]++
		deg[y]++
		addedge(x, y)
		addedge(y, x)
	}
	Tarjan(1, 0)
	ans := (POWMOD(3, m) - 3*POWMOD(2, m)%MOD + 3 + MOD) % MOD
	for i := 1; i <= tot; i++ {
		if siz[i] == 3 {
			tmp0, tmp1, tmp2, tmp3 := 0, 0, 0, 0
			if deg[qwq[i][0]] == 2 {
				tmp0 = 1
			}
			if deg[qwq[i][1]] == 2 {
				tmp1 = 1
			}
			if deg[qwq[i][2]] == 2 {
				tmp2 = 1
			}
			if tmp0+tmp1+tmp2 >= 2 {
				tmp3 = 1
			}
			ans = (ans - 6*tmp3 + MOD) % MOD
		}
	}
	for i := 1; i <= n; i++ {
		ans = (ans - POWMOD(3, cut[i]) + 3*POWMOD(2, cut[i])%MOD - 3 + MOD) % MOD
	}
	if n == 4 && m > 4 {
		ans = (ans - 6 + MOD) % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353

func POWMOD(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
