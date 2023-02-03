package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 400005

var a, b, lc, rc, stk, pos, slp, dis, ls, rs, rt, f [maxn]int
var top, tot int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		k := top
		for k != 0 && b[stk[k]] < b[i] {
			k--
		}
		if k > 0 {
			rc[stk[k]] = i
		}
		if k < top {
			lc[i] = stk[k+1]
		}
		top = k
		top++
		stk[top] = i
	}
	dfs(stk[1])
	fmt.Println(f[stk[1]])
}
func dfs(x int) {
	if lc[x] == 0 && rc[x] == 0 {
		f[x] = a[x] * b[x]
		push(&rt[x], 0, -b[x])
		push(&rt[x], a[x], b[x])
		return
	}
	if lc[x] != 0 {
		dfs(lc[x])
		f[x] += f[lc[x]]
		rt[x] = merge(rt[x], rt[lc[x]])
	}
	if rc[x] != 0 {
		dfs(rc[x])
		f[x] += f[rc[x]]
		rt[x] = merge(rt[x], rt[rc[x]])
	}
	push(&rt[x], a[x], 0)
	push(&rt[x], 0, b[x])
	sum := 0
	for pos[rt[x]] < a[x] || sum+slp[rt[x]] < 0 {
		sum += slp[rt[x]]
		rec := pos[rt[x]]
		pop(&rt[x])
		f[x] += sum * (pos[rt[x]] - rec)
	}
	push(&rt[x], pos[rt[x]], sum)
	push(&rt[x], 0, -b[x])
}

func merge(a, b int) int {
	if a == 0 || b == 0 {
		return a | b
	}
	if pos[a] > pos[b] || (pos[a] == pos[b] && slp[a] > slp[b]) || (pos[a] == pos[b] && slp[a] == slp[b] && a > b) {
		a, b = b, a
	}
	rs[a] = merge(rs[a], b)
	if dis[ls[a]] < dis[rs[a]] {
		ls[a], rs[a] = rs[a], ls[a]
	}
	dis[a] = dis[ls[a]] + 1
	return a
}

func push(a *int, p, s int) {
	tot++
	pos[tot] = p
	slp[tot] = s
	*a = merge(*a, tot)
}

func pop(a *int) {
	*a = merge(ls[*a], rs[*a])
}
